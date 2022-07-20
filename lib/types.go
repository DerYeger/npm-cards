package lib

import svg "github.com/ajstarks/svgo"

type PackageData struct {
  Name string
  WeeklyDownloads []PackageDownloads
}

type PackageDownloads struct {
  Downloads int `json:"downloads"`
  Start string `json:"start"`
  End string `json:"end"`
}

type Card struct {
  SVG *svg.SVG
  PackageData PackageData
  Size int
  Padding int
  CardSize int
  BorderRadius int
}
