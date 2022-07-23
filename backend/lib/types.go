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
  BorderRadius int
}

func (card *Card) Weeks() int {
  return len(card.PackageData.WeeklyDownloads)
}

func (card *Card) CardSize() int {
  return card.Size - 2 * card.Padding
}

func (card *Card) LeftBound() float64 {
  return float64(card.Padding)
}

func (card *Card) RightBound() float64 {
  return float64(card.Size - card.Padding)
}

func (card *Card) TopBound() float64 {
  return float64(card.Padding)
}

func (card *Card) BottomBound() float64 {
  return float64(card.Size - card.Padding)
}
