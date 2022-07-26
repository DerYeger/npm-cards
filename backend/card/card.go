package card

import (
	"fmt"
	"math"

	"github.com/DerYeger/npm-cards/backend/lib"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
  card.SVG.Roundrect(int(card.LeftBound()), int(card.TopBound()), card.CardSize(), card.CardSize(), card.BorderRadius, card.BorderRadius, fmt.Sprintf("fill:%s;stroke:none;", card.BackgroundColor()))
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

  segmentWidth := float64(card.CardSize()) / float64((card.Weeks() - 1))

  path := ""
  for i := 0; i < card.Weeks(); i++ {
    xCord := card.LeftBound() + float64(i) * segmentWidth
    availableHeight := card.CardSize() / 2 - card.BorderRadius
    yCord := float64(card.PackageData.WeeklyDownloads[i].Downloads) / float64(downloadsMax) * float64(availableHeight)
    yCord = float64(card.BottomBound()) - yCord - float64(card.BorderRadius) - strokeWidth / 2
    if i == 0 {
      // Add extra point at left edge for smooth cutoff
      path = fmt.Sprintf("M%f,%f", card.LeftBound(), yCord)
      xCord = card.LeftBound() + strokeWidth
    } else if i == card.Weeks() - 1 {
      // Add extra point at right edge for smooth cutoff
      path = fmt.Sprintf("%sL%f,%f", path, card.RightBound() - strokeWidth, yCord)
      xCord = card.RightBound()
    }
    path = fmt.Sprintf("%sL%f,%f", path, xCord, yCord)
  }

  card.SVG.Path(path, fmt.Sprintf("fill:none;stroke:url(#graph);stroke-width:%fpx;stroke-linejoin:round;", strokeWidth))
}

func makeDefs(card *lib.Card) {
  s := card.SVG
  graphGradient := card.GraphGradient()

  s.Def()
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
  titleColor := card.TitleColor()
  titleSize := math.Min(availableSpace / (0.5 * float64(len(title))), fontSize)
  s.Text(int(textStart), int(textStart + titleSize / 2), title, fmt.Sprintf("dominant-baseline:middle;color:%s;fill:%s;font-size:%fpx;font-family:sans-serif;", titleColor, titleColor, titleSize))

  recentDownloads :=  card.PackageData.WeeklyDownloads[card.Weeks() - 1].Downloads

  printer := message.NewPrinter(language.English)
  subtitle := printer.Sprintf("%d downloads last week", recentDownloads)
  subtitleColor := card.SubtitleColor()
  subtitleSize := math.Min(availableSpace / (0.5 * float64(len(subtitle))), 2.0 / 3.0 * titleSize)
  s.Text(int(textStart), int(textStart + titleSize / 2 + textPadding + subtitleSize / 2), subtitle, fmt.Sprintf("dominant-baseline:middle;color:%s;fill:%s;font-size:%fpx;font-family:sans-serif;", subtitleColor, subtitleColor, subtitleSize))
}
