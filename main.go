// json-viewer/main.go
package main

import (
	"fmt"
	"os"

	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
	"github.com/SzymonSolecki/json-log-viewer/pkg/ui"
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

	app := ui.NewApp()

	if err := app.Run(headers, parsedData); err != nil {
		panic(err)
	}
}
