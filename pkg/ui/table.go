// Package ui
package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
)

type DataTable struct {
	*tview.Table
	headers      data.Headers
	originalData data.RowData
}

func NewDataTable() *DataTable {
	table := tview.NewTable().
		SetBorders(true).SetFixed(1, 0)

	return &DataTable{
		Table: table,
	}
}

func (dt *DataTable) Populate(headers data.Headers, rowData data.RowData) {
	dt.headers = headers
	dt.originalData = rowData
	dt.UpdateView(rowData)
}

func (dt *DataTable) UpdateView(rowData data.RowData) {
	dt.Clear()

	greenColor := tcell.ColorGreen
	whiteColor := tcell.ColorWhite
	for i, value := range dt.headers {
		dt.SetCell(0, i,
			tview.NewTableCell(value).
				SetTextColor(greenColor).
				SetAlign(tview.AlignCenter).
				SetExpansion(1),
		)
	}

	for i, row := range rowData {
		for j, value := range dt.headers {

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

func (dt *DataTable) ResetFilter() {
	dt.UpdateView(dt.originalData)
}

func (dt *DataTable) ApplyFilter(filterColumnIndex int, filterValue string) {
	filteredData := data.RowData{}
	header := dt.headers[filterColumnIndex]

	for _, v := range dt.originalData {
		value, ok := v[header]
		if !ok {
			continue
		}
		cellValue := strings.ToLower(fmt.Sprintf("%s", value))
		if strings.Contains(cellValue, strings.ToLower(filterValue)) {
			filteredData = append(filteredData, v)
		}
	}

	dt.UpdateView(filteredData)
}
