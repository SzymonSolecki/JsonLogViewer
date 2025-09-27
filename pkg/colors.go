// Package pkg
package pkg

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetGlobalStyle() {
	tview.Styles.PrimitiveBackgroundColor = tcell.Color236

	tview.Styles.SecondaryTextColor = tcell.ColorSeaGreen
	tview.Styles.TertiaryTextColor = tcell.ColorMediumPurple

	tview.Styles.BorderColor = tcell.ColorBurlyWood
	tview.Styles.TitleColor = tcell.ColorBurlyWood
	tview.Styles.GraphicsColor = tcell.ColorBurlyWood
}
