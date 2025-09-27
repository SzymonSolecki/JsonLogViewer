// Package ui
package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
)

type DataTable struct {
	*tview.Table
	headers data.Headers
}

func NewDataTable() *DataTable {
	table := tview.NewTable().
		SetBorders(true).SetFixed(1, 0)

	return &DataTable{
		Table: table,
	}
}

func (dt *DataTable) Pouplate(headers data.Headers, rowData data.RowData) {
	greenColor := tcell.ColorGreen
	whiteColor := tcell.ColorWhite
	for i, value := range headers {
		dt.SetCell(0, i,
			tview.NewTableCell(value).
				SetTextColor(greenColor).
				SetAlign(tview.AlignCenter).
				SetExpansion(1),
		)
	}

	for i, row := range rowData {
		for j, value := range headers {

			formattedCell := fmt.Sprintf("%v", row[value])
			dt.SetCell(i+1, j,
				tview.NewTableCell(formattedCell).
					SetTextColor(whiteColor).
					SetAlign(tview.AlignCenter).
					SetExpansion(1),
			)
		}
	}
}
