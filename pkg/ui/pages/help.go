package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getHelpPage1() string {
	return `[seagreen]Navigation
[burlywood]Left arrow, h[white]: Move left.
[burlywood]Right arrow, l[white]: Move right.
[burlywood]Down arrow, j[white]: Move down.
[burlywood]Up arrow, k[white]: Move up.
[burlywood]0[white]: Move to the first column.
[burlywood]$[white]: Move to the last column.
[burlywood]g[white]: Move to the first row.
[burlywood]G[white]: Move to the last row.

[seagreen]Filtering
[burlywood]f[white]: Filter column by value from the selected cell.
[burlywood]Enter[white]: Filter column by provided value.

[seagreen]Other
[burlywood]y[white]: Copy value from the selected cell.
[burlywood]v[white]: Preview selected cell in bigger window.
[burlywood]Ctrl+C, q[white]: Quit the app.

[mediumpurple]Press press Escape to return.`
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
