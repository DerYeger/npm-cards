package lib

import (
	"strings"

	svg "github.com/ajstarks/svgo"
)

type PackageData struct {
  Name string
  WeeklyDownloads []PackageDownloads
}

type PackageDownloads struct {
  Downloads int `json:"downloads"`
  Start string `json:"start"`
  End string `json:"end"`
}

type Theme int64

const (
  Dark Theme = iota
  DarkTransparent
  Light
  LightTransparent
)

var (
  themeMap = map[string]Theme{
    "dark": Dark,
    "dark-transparent": DarkTransparent,
    "light": Light,
    "light-transparent": LightTransparent,
  }
)

func StringToTheme(str string) Theme {
    theme, ok := themeMap[strings.ToLower(str)]
    if !ok {
      return Dark
    }
    return theme
}

type Card struct {
  SVG *svg.SVG
  PackageData PackageData
  Size int
  Padding int
  BorderRadius int
  Theme Theme
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

func (card *Card) TitleColor() string {
  if card.Theme == Light || card.Theme == LightTransparent {
    return "#000"
  }
  return "#fff"
}

func (card *Card) SubtitleColor() string {
  if card.Theme == Light || card.Theme == LightTransparent {
    return "#222"
  }
  return "#ddd"
}

func (card *Card) BackgroundColor() string {
  if card.Theme == Light {
    return "#eee"
  } else if card.Theme == Dark {
    return "#111"
  }
  return "transparent"
}

func (card *Card) GraphGradient() []svg.Offcolor {
  if card.Theme == Light || card.Theme == LightTransparent {
    return []svg.Offcolor {
      {
        Color: "green",
        Opacity: 100,
        Offset: 0,
      },
      {
        Color: "forestgreen",
        Opacity: 100,
        Offset: 30,
      },
      {
        Color: "mediumseagreen",
        Opacity: 100,
        Offset: 60,
      },
      {
        Color: "darkorange",
        Opacity: 100,
        Offset: 80,
      },
      {
        Color: "red",
        Opacity: 100,
        Offset: 100,
      },
    }
  }
  return []svg.Offcolor {
      {
        Color: "forestgreen",
        Opacity: 100,
        Offset: 0,
      },
      {
        Color: "limegreen",
        Opacity: 100,
        Offset: 50,
      },
      {
        Color: "greenyellow",
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
}
