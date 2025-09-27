// json-viewer/main.go
package main

import (
	"fmt"
	"os"

	"github.com/SzymonSolecki/json-log-viewer/pkg"
	"github.com/SzymonSolecki/json-log-viewer/pkg/data"
	"github.com/SzymonSolecki/json-log-viewer/pkg/ui/pages"
	"github.com/rivo/tview"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Error: No data piped.")
		fmt.Println("Usage: cat data.json | json-viewer")
		os.Exit(1)
	}

	parsedData, err := data.ParseJSON(os.Stdin)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	if len(parsedData.Headers) == 0 {
		fmt.Println("No headers found")
		os.Exit(0)
	}
	if len(parsedData.Data) == 0 {
		fmt.Println("No data found")
		os.Exit(0)
	}

	pkg.SetGlobalStyle()
	app := tview.NewApplication()
	pages := pages.NewRootPage(parsedData, app)

	app.SetRoot(pages, true)

	if err := app.EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
