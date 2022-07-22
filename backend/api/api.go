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
  r := gin.Default()
  cacheStore, cacheDuration := createCacheStore()

  r.Use(gin.Logger())
  r.Use(gin.Recovery())

	r.GET("/api/packages/*name", cache.CacheByRequestURI(cacheStore, cacheDuration, cache.IgnoreQueryOrder()), handleRequest)

  r.GET("/api/health", func(ctx *gin.Context) {
    ctx.Status(http.StatusOK)
  })

  log.Printf("Listening on %d\n", port)
	r.Run(fmt.Sprintf(":%d", port))
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
