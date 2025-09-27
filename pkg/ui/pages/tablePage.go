package pages

import (
	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
	"github.com/SzymonSolecki/json-log-viewer/pkg/ui"
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewTablePage(tableData *data.ParsedData, app *tview.Application, pages *tview.Pages, viewArea *tview.TextView) *tview.Flex {
	dt := ui.NewDataTable()
	dt.Populate(tableData)
	dt.SetWrapSelection(true, false)

	tf := tview.NewInputField()
	tf.SetBorder(true)
	tf.SetPlaceholderStyle(tf.GetPlaceholderStyle().
		Background(tview.Styles.PrimitiveBackgroundColor).
		Foreground(tview.Styles.SecondaryTextColor),
	)
	tf.SetFieldBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	tf.SetFieldTextColor(tview.Styles.SecondaryTextColor)
	tf.SetPlaceholder("Press F1 for help, press Ctrl-C or q to exit")

	additionalInfo := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetWordWrap(true)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tf, 0, 5, false).
		AddItem(dt, 0, 100, true).
		AddItem(additionalInfo, 1, 0, false)

	dt.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'f':
			row, col := dt.GetSelection()
			dt.ApplyFilter(dt.GetCell(row, col).Text, additionalInfo)
			return nil
		case 'y':
			row, col := dt.GetSelection()
			clipboard.WriteAll(dt.GetCell(row, col).Text)
			return nil
		case 'v':
			row, col := dt.GetSelection()
			viewArea.SetText(dt.GetCell(row, col).Text)
			pages.SwitchToPage("view")
			return nil
		case 'g':
			_, col := dt.GetSelection()
			dt.Select(0, col)
			return nil
		case 'G':
			_, col := dt.GetSelection()
			dt.Select(dt.GetRowCount()-1, col)
			return nil
		case '0':
			row, _ := dt.GetSelection()
			dt.Select(row, 0)
			return nil
		case '$':
			row, _ := dt.GetSelection()
			dt.Select(row, dt.GetColumnCount()-1)
			return nil
		case 'q':
			app.Stop()
			return nil
		}
		if event.Key() == tcell.KeyEsc {
			dt.ResetFilter()
			additionalInfo.SetText("")
			return nil
		}
		return event
	})

	tf.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			result := tf.GetText()
			row, col := dt.GetSelection()
			if result == "" {
				result = dt.GetCell(row, col).Text
			}
			dt.ApplyFilter(result, additionalInfo)
		}
		tf.SetText("")
		app.SetFocus(dt)
	})

	dt.SetSelectedFunc(func(row int, column int) {
		app.SetFocus(tf)
	})
	return layout
}
