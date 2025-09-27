// Package pages
package pages

import (
	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewRootPage(tableData *data.ParsedData, app *tview.Application) *tview.Pages {
	root := tview.NewPages()

	helpView := NewHelpPage(root)
	viewFrame, viewArea := NewViewPage(root)
	mainView := NewTablePage(tableData, app, root, viewArea)

	root.AddAndSwitchToPage("main", mainView, true)
	root.AddPage("help", helpView, true, false)
	root.AddPage("view", viewFrame, true, false)

	root.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			root.ShowPage("help")
			return nil
		}
		return event
	})

	return root
}
