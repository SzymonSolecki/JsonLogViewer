// Package ui
package ui

import (
	"fmt"

	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app        *tview.Application
	table      *DataTable
	textField  *tview.InputField
	miscLayout *tview.Flex
	mainLayout *tview.Flex
}

func NewApp() *App {
	dt := NewDataTable()

	tf := tview.NewInputField()
	tf.SetFieldBackgroundColor(tcell.ColorBlack)
	tf.SetPlaceholderStyle(tf.GetPlaceholderStyle().Background(tcell.ColorBlack))
	tf.SetPlaceholder("Filter by...")

	a := &App{
		app:       tview.NewApplication(),
		table:     dt,
		textField: tf,
	}

	dt.SetSelectedFunc(func(row int, column int) {
		a.showFilter()
	})

	a.miscLayout = tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(a.textField, 0, 1, false)
	a.miscLayout.SetBorder(true)

	a.mainLayout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.miscLayout, 0, 5, false).
		AddItem(a.table, 0, 100, true)

	a.app.SetFocus(dt)
	a.setupKeybindings()
	return a
}

func (a *App) setupKeybindings() {
	a.table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			a.app.Stop()
			return nil
		case 'f':
			row, col := a.table.GetSelection()
			result := a.table.GetCell(row, col).Text
			a.table.ApplyFilter(result)
			return nil
		case 'y':
			return nil
		case 'v':
			return nil
		default:
			if event.Key() == tcell.KeyEsc {
				a.table.ResetFilter()
				return nil
			}
		}
		return event
	})

	a.textField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			result := a.textField.GetText()
			row, col := a.table.GetSelection()
			if result == "" {
				result = a.table.GetCell(row, col).Text
			}
			a.table.ApplyFilter(result)
		}
		a.hideFilter()
	})
}

func (a *App) showFilter() {
	a.app.SetFocus(a.textField)
}

func (a *App) hideFilter() {
	a.textField.SetText("")
	a.app.SetFocus(a.table)
}

func (a *App) Run(headers data.Headers, rowData data.RowData) error {
	a.table.Populate(headers, rowData)

	if err := a.app.SetRoot(a.mainLayout, true).EnableMouse(true).Run(); err != nil {
		return fmt.Errorf("error running application: %w", err)
	}
	return nil
}
