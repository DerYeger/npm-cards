package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	svg "github.com/ajstarks/svgo"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"

	"github.com/DerYeger/npm-cards/backend/card"
	"github.com/DerYeger/npm-cards/backend/lib"
	"github.com/DerYeger/npm-cards/backend/npm"
)

func createCacheStore() (persist.CacheStore, time.Duration) {
  redisUrl := os.Getenv("REDIS_URL")
  if redisUrl != "" {
    log.Println("Using Redis cache store with 1 day duration")
    opts, err := redis.ParseURL(redisUrl)
    if err != nil {
      log.Panic(err)
    }
    return persist.NewRedisStore(redis.NewClient(opts)), 24 * time.Hour
  }

  log.Println("Using in-memory cache store with 1 minute duration")
  return persist.NewMemoryStore(1 * time.Minute), 1 * time.Minute
}

func StartServer(port int) {
  router := gin.Default()
  cacheStore, cacheDuration := createCacheStore()

  router.Use(gin.Logger())
  router.Use(gin.Recovery())
  router.Use(cors.Default())

  router.GET("/", handleFallbackPage)
  router.GET("/api/packages", handleFallbackPage)

	router.GET("/api/packages/*name", cache.CacheByRequestURI(cacheStore, cacheDuration, cache.IgnoreQueryOrder()), handleRequest)

  router.GET("/api/health", func(ctx *gin.Context) {
    ctx.Status(http.StatusOK)
  })

  log.Printf("Listening on %d\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}

var (
  defaultWeeks = 16
  defaultSize = 512
  defaultPadding = 0
  defaultBorderRadius = 16
)

func handleRequest(ctx *gin.Context) {
  weeks, err := strconv.Atoi(ctx.Query("weeks"))
  if err != nil {
    weeks = defaultWeeks
  }
  if weeks < 2 {
    ctx.AbortWithStatus(http.StatusBadRequest)
    return
  }

  packageName := ctx.Param("name")
  if packageName != "" && packageName[0] == '/' {
    packageName = packageName[1:]
  }
  if packageName == "" {
    ctx.AbortWithStatus(http.StatusBadRequest)
    return
  }

  size, err := strconv.Atoi(ctx.Query("size"))
  if err != nil {
    size = defaultSize
  }

  padding, err := strconv.Atoi(ctx.Query("padding"))
  if err != nil {
    padding = defaultPadding
  }

  borderRadius, err := strconv.Atoi(ctx.Query("borderRadius"))
  if err != nil {
    borderRadius = defaultBorderRadius
  }

  packageData, err := npm.GetPackageData(packageName, weeks)
  if err != nil {
    statusCode := err.(*npm.ApiError).StatusCode
    if http.StatusText(statusCode) != "" {
      ctx.AbortWithStatus(statusCode)
      return
    }
    panic(err)
  }

  ctx.Header("Content-Type", "image/svg+xml")
  ctx.Header("Cache-Control", "public, max-age=86400, immutable")

  cardData := lib.Card {
    SVG: svg.New(ctx.Writer),
    PackageData: packageData,
    Size: size,
    Padding: padding,
    BorderRadius: borderRadius,
  }

  card.CreateCard(&cardData)
}

func handleFallbackPage(ctx *gin.Context) {
  ctx.Header("Content-Type", "text/html; charset=utf-8")
  ctx.Header("Cache-Control", "public, max-age=86400, immutable")
  ctx.String(http.StatusOK, `
    <html>
      <body>
        <h1>NPM Cards API</h1>
        <p>
          Welcome to NPM Cards! The API is available at <a href="/api/packages/:packageName">/api/packages/:packageName</a>.
        </p>
        <h2>Examples:</h2>
        <div>
          <a href="/api/packages/react?size=256&padding=0&borderRadius=16&weeks=64">
            <img alt="React" src="/api/packages/react?size=256&padding=0&borderRadius=16&weeks=64">
          </a>
          <a href="/api/packages/vite?size=256&padding=0&borderRadius=16&weeks=64">
            <img alt="Vite" src="/api/packages/vite?size=256&padding=0&borderRadius=16&weeks=64">
          </a>
          <a href="/api/packages/@yeger/vue-masonry-wall?size=256&padding=0&borderRadius=16&weeks=64">
            <img alt="@yeger/vue-masonry-wall" src="/api/packages/@yeger/vue-masonry-wall?size=256&padding=0&borderRadius=16&weeks=64">
          </a>
          </ul>
        </div>
      </body>
    </html>`,
  )
}
