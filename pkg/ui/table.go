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

func (dt *DataTable) ApplyFilter(filterText string) {
	// If the filter is empty, show all data
	if strings.TrimSpace(filterText) == "" {
		dt.UpdateView(dt.originalData)
		return
	}

	parts := strings.SplitN(filterText, ":", 2)
	if len(parts) != 2 {
		return // Invalid filter format
	}

	colIndexStr, searchTerm := parts[0], strings.ToLower(strings.TrimSpace(parts[1]))
	var colIndex int
	if _, err := fmt.Sscanf(colIndexStr, "%d", &colIndex); err != nil {
		return // Invalid column index
	}

	if colIndex < 0 || colIndex >= len(dt.headers) {
		return // Column index out of bounds
	}

	headerToFilter := dt.headers[colIndex]
	filteredData := make([]map[string]any, 0)

	for _, row := range dt.originalData {
		value, ok := row[headerToFilter]
		if !ok {
			continue
		}
		cellValue := strings.ToLower(fmt.Sprintf("%v", value))
		if strings.Contains(cellValue, searchTerm) {
			filteredData = append(filteredData, row)
		}
	}

	dt.UpdateView(filteredData)
}
