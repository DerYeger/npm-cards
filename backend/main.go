package main

import (
	"log"
	"os"
	"strconv"

	"github.com/DerYeger/npm-cards/backend/api"
)

func main() {
  port, err := strconv.Atoi(os.Getenv("PORT"))
  if err != nil {
    log.Print("Invalid or missing port. Defaulting to 8080.")
    port = 8080
  }
  api.StartServer(port)
}
