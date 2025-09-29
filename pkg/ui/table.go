// Package ui
package ui

import (
	"fmt"
	"strings"

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
		cell := tview.NewTableCell(value).
			SetTextColor(tview.Styles.SecondaryTextColor).
			SetAlign(tview.AlignCenter).
			SetExpansion(1)
		cell.SetStyle(cell.Style.Bold(true))
		dt.SetCell(0, i, cell)
	}

	for i, row := range rowData {
		for j, value := range dt.headers {

			formattedCell := fmt.Sprintf("%v", row[value])
			dt.SetCell(i+1, j,
				tview.NewTableCell(formattedCell).
					SetAlign(tview.AlignCenter).
					SetExpansion(1).
					SetMaxWidth(70),
			)
		}
	}
}

func (dt *DataTable) ResetFilter() {
	dt.UpdateView(dt.originalData)
}

func (dt *DataTable) ApplyFilter(filterValue string, statusText ...*tview.TextView) {
	filteredData := data.RowData{}
	_, col := dt.GetSelection()
	header := dt.headers[col]

	for _, row := range dt.originalData {
		value := row[header]
		cellValue := strings.ToLower(fmt.Sprintf("%v", value))
		if strings.Contains(cellValue, strings.ToLower(filterValue)) {
			filteredData = append(filteredData, row)
		}
	}

	dt.UpdateView(filteredData)

	if len(statusText) > 0 {
		item := statusText[0]
		tmp := fmt.Sprintf("Current filter - [burlywood]Column: [seagreen]%v [white]| [burlywood]Value: [seagreen]%v", dt.headers[col], filterValue)
		item.SetText(tmp)
	}
}
