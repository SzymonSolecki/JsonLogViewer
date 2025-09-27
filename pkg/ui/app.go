// Package ui
package ui

import (
	"fmt"

	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	tviewApp *tview.Application
	table    *DataTable
	display  *tview.InputField
	miscBox  *tview.Flex
	layout   *tview.Flex

	selectedRow int
	selectedCol int
}

func NewApp() *App {
	a := &App{
		tviewApp: tview.NewApplication(),
		table:    NewDataTable(),
	}

	a.display = tview.NewInputField()
	a.display.SetFieldBackgroundColor(tcell.ColorBlack)

	a.miscBox = tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(a.display, 0, 1, false)
	a.miscBox.SetBorder(true)

	a.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.miscBox, 0, 5, true).
		AddItem(a.table, 0, 100, true)

	a.setupKeybindings()
	return a
}

func (a *App) setupKeybindings() {
	a.tviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			a.tviewApp.Stop()
			return nil
		case '/':
			a.tviewApp.SetFocus(a.table)
			a.table.Select(0, 0)
			a.table.SetSelectable(true, true)
			a.table.SetSelectedFunc(func(row int, column int) {
				a.table.SetSelectable(false, false)
				a.selectedCol = column
				a.selectedRow = row
				a.showFilter()
			})
			return nil
		default:
			if event.Key() == tcell.KeyEsc {
				a.table.SetSelectable(false, false)
				a.table.ResetFilter()
			}
		}
		return event
	})
	a.display.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			result := a.display.GetText()
			if result == "" {
				result = a.table.GetCell(a.selectedRow, a.selectedCol).Text
			}
			a.table.ApplyFilter(a.selectedCol, result)
		}
		a.hideFilter()
	})
}

func (a *App) showFilter() {
	a.tviewApp.SetFocus(a.display)
}

func (a *App) hideFilter() {
	a.display.SetText("")
	a.tviewApp.SetFocus(a.table)
}

func (a *App) Run(headers data.Headers, rowData data.RowData) error {
	a.table.Populate(headers, rowData)

	if err := a.tviewApp.SetRoot(a.layout, true).EnableMouse(true).Run(); err != nil {
		return fmt.Errorf("error running application: %w", err)
	}
	return nil
}
