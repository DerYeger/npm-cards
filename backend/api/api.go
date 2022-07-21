package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	svg "github.com/ajstarks/svgo"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"

	"github.com/DerYeger/npm-cards/backend/card"
	"github.com/DerYeger/npm-cards/backend/lib"
	"github.com/DerYeger/npm-cards/backend/npm"
)

func StartServer(port int) {
  log.Printf("Listening on %d", port)
  r := gin.Default()
  memoryStore := persist.NewMemoryStore(1 * time.Minute)
  r.Use(gin.Logger())
  r.Use(gin.Recovery())
	r.GET("/api/packages/*name", cache.CacheByRequestURI(memoryStore, 1 * time.Minute), handleRequest)
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
    log.Print(err)
    statusCode := err.(*npm.ApiError).StatusCode
    log.Print(statusCode)
    if http.StatusText(statusCode) != "" {
      ctx.AbortWithStatus(statusCode)
    return
  }
    panic(err)
  }

  ctx.Header("Content-Type", "image/svg+xml")

  cardData := lib.Card {
    SVG: svg.New(ctx.Writer),
    PackageData: packageData,
    Size: size,
    Padding: padding,
    CardSize: size - 2 * padding,
    BorderRadius: borderRadius,
  }

  card.CreateCard(cardData)
}
