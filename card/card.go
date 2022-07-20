package card

import (
	"fmt"

	"github.com/DerYeger/npm-cards/lib"
	svg "github.com/ajstarks/svgo"
)

func CreateCard(card lib.Card) {
  s := card.SVG
  s.Start(card.Size, card.Size)

  makeDefs(card)

  makeBackground(card)
  makeGraph(card)
  makeText(card)

  s.End()
}

func makeBackground(card lib.Card) {
  card.SVG.Roundrect(card.Padding, card.Padding, card.CardSize, card.CardSize, card.BorderRadius, card.BorderRadius, "fill:url(#background);stroke:none;")
}

func calculateDownloadPoint(card lib.Card, downloads int, downloadsMax int) int {
  availableHeight := (card.Size - 2 * card.Padding) / 2 - card.BorderRadius
  point := int(float64(downloads) / float64(downloadsMax) * float64(availableHeight))
  point = (card.Size - 2 * card.Padding) - point - card.BorderRadius
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

  path := "M" + fmt.Sprint(card.Padding) + "," + fmt.Sprint(calculateDownloadPoint(card, card.PackageData.WeeklyDownloads[0].Downloads, downloadsMax))
  for i := 1; i < len(card.PackageData.WeeklyDownloads); i++ {
    xCord := card.Padding + i * segmentWidth
    if i == len(card.PackageData.WeeklyDownloads) - 1 {
      xCord = card.Size - card.Padding
    }
    path = path + "L" + fmt.Sprint(xCord)  + "," + fmt.Sprint(calculateDownloadPoint(card, card.PackageData.WeeklyDownloads[i].Downloads, downloadsMax))
  }
  card.SVG.Path(path, "fill:none;stroke:url(#graph);stroke-width:5px;stroke-linejoin:round;")
}

func makeDefs(card lib.Card) {
  s := card.SVG
  backgroundGradient := []svg.Offcolor {
    {
      Color: "rgba(0, 0, 0, 0.8)",
      Opacity: 100,
      Offset: 0,
    },
    {
      Color: "rgba(0, 0, 0, 0.9)",
      Opacity: 100,
      Offset: 100,
    },
  }

  graphGradient := []svg.Offcolor {
    {
      Color: "green",
      Opacity: 100,
      Offset: 0,
    },
    {
      Color: "greenyellow",
      Opacity: 100,
      Offset: 50,
    },
    {
      Color: "yellow",
      Opacity: 100,
      Offset: 65,
    },
    {
      Color: "orange",
      Opacity: 100,
      Offset: 80,
    },
    {
      Color: "red",
      Opacity: 100,
      Offset: 100,
    },
  }

  s.Def()
  s.LinearGradient("background", 50, 0, 50, 200, backgroundGradient)
  s.LinearGradient("graph", 50, 0, 50, 200, graphGradient)
  s.DefEnd()
}

func makeText(card lib.Card) {
  s := card.SVG

  textPadding := 32
  textStart := card.Padding + textPadding

  s.Text(textStart, textStart + 32 / 2, card.PackageData.Name, "dominant-baseline:middle;fill:white;font-size:32px;font-family:sans-serif;")

  s.Text(textStart, textStart + 32 / 2 + textPadding + 24 / 2, fmt.Sprint(card.PackageData.WeeklyDownloads[len(card.PackageData.WeeklyDownloads) - 1].Downloads) + " downloads last week", "dominant-baseline:middle;fill:darkgray;font-size:24px;font-family:sans-serif;")
}
