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
		SetBorders(true).SetFixed(1, 0).SetSelectable(true, true)

	return &DataTable{
		Table: table,
	}
}

func (dt *DataTable) Populate(tableData *data.ParsedData) {
	dt.headers = tableData.Headers
	dt.originalData = tableData.Data

	dt.UpdateView(dt.originalData)
}

func (dt *DataTable) UpdateView(rowData data.RowData) {
	dt.Clear()

	for i, value := range dt.headers {
		dt.SetCell(0, i,
			tview.NewTableCell(value).
				SetTextColor(tcell.ColorGreen).
				SetAlign(tview.AlignCenter).
				SetExpansion(1),
		)
	}

	for i, row := range rowData {
		for j, value := range dt.headers {

			formattedCell := fmt.Sprintf("%v", row[value])
			dt.SetCell(i+1, j,
				tview.NewTableCell(formattedCell).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignCenter).
					SetExpansion(1),
			)
		}
	}
}

func (dt *DataTable) ResetFilter() {
	dt.UpdateView(dt.originalData)
}

func (dt *DataTable) ApplyFilter(filterValue string) {
	filteredData := data.RowData{}
	_, col := dt.GetSelection()
	header := dt.headers[col]

	for _, v := range dt.originalData {
		value, ok := v[header]
		if !ok {
			continue
		}
		cellValue := strings.ToLower(fmt.Sprintf("%v", value))
		if strings.Contains(cellValue, strings.ToLower(filterValue)) {
			filteredData = append(filteredData, v)
		}
	}

	dt.UpdateView(filteredData)
}
