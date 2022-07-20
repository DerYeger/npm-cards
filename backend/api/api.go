package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	svg "github.com/ajstarks/svgo"

	"github.com/DerYeger/npm-cards/backend/card"
	"github.com/DerYeger/npm-cards/backend/lib"
	"github.com/DerYeger/npm-cards/backend/npm"
)

func StartServer(port int) {
  http.Handle("/", http.HandlerFunc(handleRequest))
  log.Printf("Listening on %d", port)
  err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
  query := req.URL.Query()

  weeks, err := strconv.Atoi(query.Get("weeks"))
  if err != nil {
    weeks = 10
  }

  packageName := query.Get("package")
  if packageName == "" {
    w.WriteHeader(400)
    return
  }

  packageData, err := npm.GetPackageData(packageName, weeks)
  if err != nil {
    w.WriteHeader(400)
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

  w.Header().Set("Content-Type", "image/svg+xml")

  cardData := lib.Card {
    SVG: svg.New(w),
    PackageData: packageData,
    Size: size,
    Padding: padding,
    CardSize: size - 2 * padding,
    BorderRadius: borderRadius,
  }

  card.CreateCard(cardData)
}
