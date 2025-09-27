package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewViewPage(pages *tview.Pages) (*tview.Frame, *tview.TextView) {
	view := tview.NewTextView().
		SetDynamicColors(true).SetScrollable(true)

	frame := tview.NewFrame(view)
	frame.SetBorder(true).
		SetTitle("Preview").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.SwitchToPage("main")
				return nil
			}
			return event
		})
	return frame, view
}
