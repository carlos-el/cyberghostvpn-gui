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

func setConnectionLabel(statusLabel *widget.Label, c *models.Country, server string) {
	if server == "" {
		statusLabel.SetText("Disconnected")
	} else if c == nil {
		statusLabel.SetText("Connected to: " + server)
	} else {
		statusLabel.SetText("Connected to: " + server + " - " + c.String())
	}
}

func disconnect(statusButton *widget.Button, statusLabel *widget.Label) {
	commander.Disconnect()
	statusButton.Disable()
	setConnectionLabel(statusLabel, nil, "")
}

func connect(statusButton *widget.Button, statusLabel *widget.Label, c *models.Country) {
	server, _ := commander.Connect(c)
	statusButton.Enable()
	setConnectionLabel(statusLabel, c, server)
}

func main() {
	a := app.New()
	myWindow := a.NewWindow("CyberGhost GUI")
	myWindow.Resize(fyne.NewSize(400, 600))

	// Status component
	statusButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {})
	statusButton.Disable()
	statusLabel := widget.NewLabel("Disconnected")
	statusButton.OnTapped = func() {
		disconnect(statusButton, statusLabel)
	}
	statusComponent := container.NewHBox(
		statusButton,
		statusLabel,
	)

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
			b := widget.NewButtonWithIcon("", theme.LoginIcon(), func() {})
			return container.NewHBox(widget.NewLabel("template"), layout.NewSpacer(), b)
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			c, _ := i.(binding.Untyped).Get()
			country := c.(models.Country)
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(country.String())
			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				connect(statusButton, statusLabel, &country)
			}
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
