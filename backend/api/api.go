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

func handleRequest(c *gin.Context) {
  weeks, err := strconv.Atoi(c.Query("weeks"))
  if err != nil {
    weeks = defaultWeeks
  }
  if weeks < 2 {
    c.Status(http.StatusBadRequest)
    return
  }

  packageName := c.Param("name")
  if packageName != "" && packageName[0] == '/' {
    packageName = packageName[1:]
  }
  if packageName == "" {
    c.Status(http.StatusBadRequest)
    return
  }

  size, err := strconv.Atoi(c.Query("size"))
  if err != nil {
    size = defaultSize
  }

  padding, err := strconv.Atoi(c.Query("padding"))
  if err != nil {
    padding = defaultPadding
  }

  borderRadius, err := strconv.Atoi(c.Query("borderRadius"))
  if err != nil {
    borderRadius = defaultBorderRadius
  }

  packageData, err := npm.GetPackageData(packageName, weeks)
  if err != nil {
    c.Status(http.StatusNotFound)
    return
  }

  c.Header("Content-Type", "image/svg+xml")

  cardData := lib.Card {
    SVG: svg.New(c.Writer),
    PackageData: packageData,
    Size: size,
    Padding: padding,
    CardSize: size - 2 * padding,
    BorderRadius: borderRadius,
  }

  card.CreateCard(cardData)
}
