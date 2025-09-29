// Package data
package data

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
)

type (
	RowData []map[string]any
	Headers []string
)

type ParsedData struct {
	Headers Headers
	Data    RowData
}

func ParseJSON(reader *os.File) (*ParsedData, error) {
	var data RowData

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err != nil {
		return &ParsedData{}, fmt.Errorf("failed to decode json: %w", err)
	}

	if len(data) == 0 {
		return &ParsedData{}, nil
	}

	headersSet := map[string]struct{}{}
	for _, row := range data {
		for header := range row {
			headersSet[header] = struct{}{}
		}
	}

	headers := slices.Collect(maps.Keys(headersSet))
	sort.Strings(headers)

	return &ParsedData{
		Headers: headers,
		Data:    data,
	}, nil
}
