package main

import (
	"fmt"
	"store/pkg/dc"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "github.com/gookit/goutil/dump"
)

func main() {
	a := app.NewWithID("guiadmin")

	win := a.NewWindow("Gui admin")
	win.Resize(fyne.NewSize(400, 400))
	tabs := container.NewAppTabs(
		container.NewTabItem("Customers", table()),
		container.NewTabItem("Carts", widget.NewLabel("World!")),
		container.NewTabItem("Cart Items", widget.NewLabel("World!")),
		container.NewTabItem("Products", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	win.SetContent(tabs)
	win.Show()
	a.Run()
}

func table() fyne.CanvasObject {
	dc := dc.NewDc(".env")
	repo := dc.GeneralRepo
	query := fmt.Sprintf(`SELECT * FROM %s;`, "products")
	resultsMap, err := repo.QueryToMap(query)
	if err != nil {
		panic(err)
	}

	var data [][]string
	// Add table header.
	data = append(data, resultsMap.ColumnNames)
	// Add table data.
	for _, row := range resultsMap.DataMap {
		var rowData []string
		for _, column := range resultsMap.ColumnNames {
			rowData = append(rowData, fmt.Sprintf("%v", row[column]))
		}
		data = append(data, rowData)
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	table.SetColumnWidth(0,100)
	table.SetColumnWidth(1,30)
	return table
}
