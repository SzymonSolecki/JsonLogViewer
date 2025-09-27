package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getHelpPage1() string {
	return `[green]Navigation

[yellow]Left arrow, h[white]: Move left.
[yellow]Right arrow, l[white]: Move right.
[yellow]Down arrow, j[white]: Move down.
[yellow]Up arrow, k[white]: Move up.
[yellow]f[white]: Filter column by value from the selected cell.
[yellow]Enter[white]: Filter column by provided value.
[yellow]y[white]: Copy value from the selected cell.
[yellow]v[white]: Preview selected cell in bigger window.

[blue]Press press Escape to return.`
}

func NewHelpPage(pages *tview.Pages) *tview.Frame {
	help1 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(getHelpPage1())
	helpFrame := tview.NewFrame(help1)
	helpFrame.SetBorder(true).
		SetTitle("Help").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			key := event.Key()
			if key == tcell.KeyEscape {
				pages.SwitchToPage("main")
				return nil
			}
			return event
		})
	return helpFrame
}
