// Package data
package data

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type (
	RowData []map[string]any
	Headers []string
)

func ParseJSON(reader *os.File) (Headers, RowData, error) {
	var data RowData

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err != nil {
		return nil, nil, fmt.Errorf("failed to decode json: %w", err)
	}

	if len(data) == 0 {
		return Headers{}, RowData{}, nil
	}

	headers := make(Headers, 0, len(data[0]))
	for key := range data[0] {
		headers = append(headers, key)
	}
	sort.Strings(headers)

	return headers, data, nil
}
