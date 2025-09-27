// json-viewer/main.go
package main

import (
	"fmt"
	"os"

	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
	"github.com/SzymonSolecki/json-log-viewer/pkg/ui"
	"github.com/rivo/tview"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Error: No data piped.")
		fmt.Println("Usage: cat data.json | json-viewer")
		os.Exit(1)
	}

	headers, parsedData, err := data.ParseJSON(os.Stdin)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	if len(headers) == 0 {
		fmt.Println("No headers found")
		os.Exit(0)
	}
	if len(parsedData) == 0 {
		fmt.Println("No data found")
		os.Exit(0)
	}
	app := tview.NewApplication()

	topBox := tview.NewBox().SetBorder(true).SetTitle("Top")
	table := ui.NewDataTable()

	table.Pouplate(headers, parsedData)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topBox, 0, 10, false).
		AddItem(table, 0, 100, true)

	if err := app.SetRoot(flex, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
