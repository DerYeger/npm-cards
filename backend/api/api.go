package api

import (
	"fmt"
	"log"
	"strconv"

	svg "github.com/ajstarks/svgo"
	"github.com/gin-gonic/gin"

	"github.com/DerYeger/npm-cards/backend/card"
	"github.com/DerYeger/npm-cards/backend/lib"
	"github.com/DerYeger/npm-cards/backend/npm"
)

func StartServer(port int) {
  log.Printf("Listening on %d", port)
  r := gin.Default()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())
	r.GET("/", handleRequest)
	r.Run(fmt.Sprintf(":%d", port))
}

func handleRequest(c *gin.Context) {
  query := c.Request.URL.Query()

  weeks, err := strconv.Atoi(query.Get("weeks"))
  if err != nil {
    weeks = 10
  }
  if weeks < 2 {
    c.Status(400)
    return
  }

  packageName := query.Get("package")
  if packageName == "" {
    c.Status(400)
    return
  }

  packageData, err := npm.GetPackageData(packageName, weeks)
  if err != nil {
    c.Status(400)
    return
  }

  size, err := strconv.Atoi(query.Get("size"))
  if err != nil {
    size = 500
  }

  padding, err := strconv.Atoi(query.Get("padding"))
  if err != nil {
    padding = 0
  }

  borderRadius, err := strconv.Atoi(query.Get("borderRadius"))
  if err != nil {
    borderRadius = 0
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
