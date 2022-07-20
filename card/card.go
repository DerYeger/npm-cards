package card

import (
	"fmt"

	"github.com/DerYeger/npm-cards/lib"
)

func CreateCard(card lib.Card) {
  s := card.SVG
  s.Start(card.Size, card.Size)

  makeBackground(card)

  makeGraph(card)

  s.Text(48, 48 + 32 - 8, card.PackageData.Name, "fill:white;font-size:32px;font-family:sans-serif;")

  // s.Grid(0, 0, 500, 500, 16, "stroke:black;")
  s.End()
}

func makeBackground(card lib.Card) {
  card.SVG.Roundrect(card.Padding, card.Padding, card.CardSize, card.CardSize, 0, 0, "fill:rgba(0, 0, 0, 0.8);stroke:black;")
}

func calculateDownloadPoint(size int, padding int, downloads int, downloadsMax int) int {
  availableHeight := (size - 2 * padding) / 2
  point := int(float64(downloads) / float64(downloadsMax) * float64(availableHeight))
  point = (size - 2 * padding) - point
  return point
}

func makeGraph(card lib.Card) {
  downloadsMax := 1
  for i := 0; i < len(card.PackageData.WeeklyDownloads); i++ {
    downloads := card.PackageData.WeeklyDownloads[i].Downloads
    if downloads > downloadsMax {
      downloadsMax = downloads
    }
  }

  segmentWidth := (card.Size - 2 * card.Padding) / (len(card.PackageData.WeeklyDownloads) - 1)

  path := "M" + fmt.Sprint(card.Padding) + "," + fmt.Sprint(calculateDownloadPoint(card.Size, card.Padding, card.PackageData.WeeklyDownloads[0].Downloads, downloadsMax))
  for i := 1; i < len(card.PackageData.WeeklyDownloads); i++ {
    xCord := card.Padding + i * segmentWidth
    if i == len(card.PackageData.WeeklyDownloads) - 1 {
      xCord = card.Size - card.Padding
    }
    path = path + "L" + fmt.Sprint(xCord)  + "," + fmt.Sprint(calculateDownloadPoint(card.Size, card.Padding, card.PackageData.WeeklyDownloads[i].Downloads, downloadsMax))
  }
  card.SVG.Path(path, "fill:none;stroke:green;stroke-width:4px;")
}
