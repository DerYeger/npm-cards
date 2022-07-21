package card

import (
	"fmt"
	"math"

	"github.com/DerYeger/npm-cards/backend/lib"
	svg "github.com/ajstarks/svgo"
)

func CreateCard(card *lib.Card) {
  s := card.SVG
  s.Start(card.Size, card.Size)

  makeDefs(card)

  makeBackground(card)
  makeGraph(card)
  makeText(card)

  s.End()
}

func makeBackground(card *lib.Card) {
  card.SVG.Roundrect(card.LeftBound(), card.TopBound(), card.CardSize(), card.CardSize(), card.BorderRadius, card.BorderRadius, "fill:url(#background);stroke:none;")
}

func makeGraph(card *lib.Card) {
  downloadsMax := 1
  for i := 0; i < card.Weeks(); i++ {
    downloads := card.PackageData.WeeklyDownloads[i].Downloads
    if downloads > downloadsMax {
      downloadsMax = downloads
    }
  }

  strokeWidth := float64(card.CardSize()) / 128

  segmentWidth := card.CardSize() / (card.Weeks() - 1)

  path := ""
  for i := 0; i < card.Weeks(); i++ {
    xCord := card.LeftBound() + i * segmentWidth
    availableHeight := card.CardSize() / 2 - card.BorderRadius
    yCord := float64(card.PackageData.WeeklyDownloads[i].Downloads) / float64(downloadsMax) * float64(availableHeight)
    yCord = float64(card.BottomBound()) - yCord - float64(card.BorderRadius) - strokeWidth / 2
    if i == 0 {
      // Add extra point at left edge for smooth cutoff
      path = "M" + fmt.Sprint(card.LeftBound()) + "," + fmt.Sprint(yCord)
      xCord = card.LeftBound() + int(strokeWidth)
    } else if i == card.Weeks() - 1 {
      // Add extra point at right edge for smooth cutoff
      path = path + "L" + fmt.Sprint(card.RightBound() - int(strokeWidth))  + "," + fmt.Sprint(yCord)
      xCord = card.RightBound()
    }
    path = path + "L" + fmt.Sprint(xCord)  + "," + fmt.Sprint(yCord)
  }



  card.SVG.Path(path, fmt.Sprintf("fill:none;stroke:url(#graph);stroke-width:%fpx;stroke-linejoin:round;", strokeWidth))
}

func makeDefs(card *lib.Card) {
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

func makeText(card *lib.Card) {
  s := card.SVG

  textPadding := float64(card.CardSize()) / 16 + math.Sqrt(float64(card.BorderRadius))
  textStart := float64(card.LeftBound()) + textPadding
  availableSpace := float64(card.CardSize()) - 2 * textPadding
  fontSize := availableSpace / 10

  title := card.PackageData.Name
  titleColor := "#fff"
  titleSize := availableSpace / (0.5 * float64(len(title)))
  if titleSize > fontSize {
    titleSize = fontSize
  }
  s.Text(int(textStart), int(textStart + titleSize / 2), title, fmt.Sprintf("dominant-baseline:middle;color:%s;fill:%s;font-size:%fpx;font-family:sans-serif;", titleColor, titleColor, titleSize))

  recentDownloads :=  card.PackageData.WeeklyDownloads[card.Weeks() - 1].Downloads
  subtitle := fmt.Sprintf("%d downloads last week", recentDownloads)
  subtitleColor := "#ccc"
  subtitleSize := availableSpace / (0.5 * float64(len(subtitle)))
  if subtitleSize > titleSize {
    subtitleSize = titleSize
  }
  s.Text(int(textStart), int(textStart + titleSize / 2 + textPadding + subtitleSize / 2), subtitle, fmt.Sprintf("dominant-baseline:middle;color:%s;fill:%s;font-size:%fpx;font-family:sans-serif;", subtitleColor, subtitleColor, subtitleSize))
}
