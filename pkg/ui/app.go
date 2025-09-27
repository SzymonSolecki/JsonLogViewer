// Package ui
package ui

import (
	"fmt"
	"strings"

	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// App encapsulates the entire TUI application.
type App struct {
	tviewApp *tview.Application
	table    *DataTable
	filter   *tview.InputField
	layout   *tview.Flex
}

// NewApp creates and initializes the application UI components.
func NewApp() *App {
	a := &App{
		tviewApp: tview.NewApplication(),
		table:    NewDataTable(),
	}

	// Create the filter input field
	a.filter = tview.NewInputField().
		SetLabelColor(tcell.ColorAqua).
		SetFieldBackgroundColor(tcell.ColorBlueViolet)

	// Create the main layout
	a.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.table, 0, 1, true).  // Table takes up all available space
		AddItem(a.filter, 1, 1, false) // Filter is initially hidden

	a.setupKeybindings()
	return a
}

// setupKeybindings configures global and component-specific key listeners.
func (a *App) setupKeybindings() {
	// Keybinding for '/' to activate filtering
	a.tviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			a.tviewApp.Stop()
			return nil
		case '/':
			a.showFilter()
			return nil
		}
		return event
	})

	// Keybindings for the filter input field (Enter/Escape)
	a.filter.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			a.table.ApplyFilter(a.filter.GetText())
		}
		a.hideFilter()
	})
}

// showFilter makes the filter input visible and focuses it.
func (a *App) showFilter() {
	headersWithIndex := make([]string, len(a.table.headers))
	for i, h := range a.table.headers {
		headersWithIndex[i] = fmt.Sprintf("%d:%s", i, h)
	}
	label := fmt.Sprintf("Filter by (%s) > ", strings.Join(headersWithIndex, ", "))
	a.filter.SetLabel(label)

	a.layout.ResizeItem(a.filter, 1, 0) // Make filter visible
	a.tviewApp.SetFocus(a.filter)
}

// hideFilter hides the filter input and returns focus to the table.
func (a *App) hideFilter() {
	a.filter.SetText("")
	a.layout.ResizeItem(a.filter, 0, 0) // Hide filter
	a.tviewApp.SetFocus(a.table)
}

func (a *App) Run(headers data.Headers, rowData data.RowData) error {
	a.table.Populate(headers, rowData)

	if err := a.tviewApp.SetRoot(a.layout, true).EnableMouse(true).Run(); err != nil {
		return fmt.Errorf("error running application: %w", err)
	}
	return nil
}
