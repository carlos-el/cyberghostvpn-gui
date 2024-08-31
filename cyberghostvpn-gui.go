package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/carlos-el/cyberghostvpn-gui/commander"
	"github.com/carlos-el/cyberghostvpn-gui/models"
)

func main() {
	a := app.New()
	myWindow := a.NewWindow("CyberGhost GUI")
	myWindow.Resize(fyne.NewSize(400, 600))

	// Get data
	var countries, _ = commander.GetCountryList()
	countryList := binding.NewUntypedList()
	for _, t := range countries {
		countryList.Append(t)
	}

	// Service list component
	list := widget.NewListWithData(
		countryList,
		func() fyne.CanvasObject {
			b := widget.NewButton("", func() {
			})
			b.SetIcon(theme.LoginIcon())
			b.Disable()
			return container.NewHBox(widget.NewLabel("template"), layout.NewSpacer(), b)
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			country, _ := i.(binding.Untyped).Get()
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(country.(models.Country).Name)
		},
	)

	listTitle := widget.NewLabel("List of countries:")
	serviceListComponent := container.NewBorder(
		listTitle, // Top
		nil,       // Bottom
		nil,       // Left
		nil,       // Right
		list,      // Rest
	)

	// Status component
	statusButton := widget.NewButton("", func() {
	})
	statusButton.SetIcon(theme.CancelIcon())
	statusButton.Disable()

	statusText := widget.NewLabel("Disconnected")

	statusComponent := container.NewHBox(
		statusButton,
		statusText,
	)

	// Main component
	mainComp := container.NewBorder(
		statusComponent,      // Top
		nil,                  // Bottom
		nil,                  // Left
		nil,                  // Right
		serviceListComponent, // Rest
	)

	myWindow.SetContent(mainComp)

	myWindow.ShowAndRun()
}
