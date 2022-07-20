package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	svg "github.com/ajstarks/svgo"

	"github.com/DerYeger/npm-cards/card"
	"github.com/DerYeger/npm-cards/lib"

	npm "github.com/DerYeger/npm-cards/npm"
)

func main() {
  port, err := strconv.Atoi(os.Getenv("PORT"))
  if err != nil {
    log.Print("Invalid or missing port. Defaulting to 8080.")
    port = 8080
  }
	http.Handle("/", http.HandlerFunc(handleRequest))
	err = http.ListenAndServe(":" + fmt.Sprint(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
  query := req.URL.Query()

  packageName := query.Get("package")
  if packageName == "" {
    w.WriteHeader(400)
    return
  }

  packageData, err := npm.GetPackageData(packageName, 20)
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

  w.Header().Set("Content-Type", "image/svg+xml")

  cardData := lib.Card {
    SVG: svg.New(w),
    PackageData: packageData,
    Size: size,
    Padding: padding,
    CardSize: size - 2 * padding,
  }

  card.CreateCard(cardData)
}
