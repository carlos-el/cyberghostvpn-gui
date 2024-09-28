package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/carlos-el/cyberghostvpn-gui/commander"
	"github.com/carlos-el/cyberghostvpn-gui/components"
	"github.com/carlos-el/cyberghostvpn-gui/models"
)

func disconnect(loader *dialog.CustomDialog) error {
	loader.Show()
	defer loader.Hide()

	err := commander.Disconnect()
	if err != nil {
		return fmt.Errorf("could not disconnect: %w", err)
	}

	return nil
}

func connect(status *components.ConnectionStatus, loader *dialog.CustomDialog, c *models.Country) error {
	loader.Show()
	defer loader.Hide()

	server, err := commander.Connect(c)
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}

	status.SetConnected(c, server)
	return nil
}

func main() {
	a := app.New()
	myWindow := a.NewWindow("CyberGhost GUI")
	myWindow.Resize(fyne.NewSize(400, 600))

	// Error dialog
	errorDialog := components.NewErrorDialog(&a, &myWindow)
	// Loader dialog to block interaction
	loader := dialog.NewCustomWithoutButtons("", widget.NewLabel("Loading..."), myWindow)
	// Status component
	connectionStatusComponent := components.NewConnectionStatus(
		func() { disconnect(loader) },
	)

	// Get data
	var countries, err = commander.GetCountryList()
	if err != nil {
		errorDialog.Show(err)
	}
	countryList := binding.NewUntypedList()
	for _, t := range countries {
		countryList.Append(t)
	}

	// Service list component
	list := widget.NewListWithData(
		countryList,
		func() fyne.CanvasObject {
			b := widget.NewButtonWithIcon("", theme.LoginIcon(), func() {
				// Defined later
			})
			return container.NewHBox(widget.NewLabel("template"), layout.NewSpacer(), b)
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			c, errCasting := i.(binding.Untyped).Get()
			if errCasting != nil {
				errorDialog.Show(errCasting)
			}
			country := c.(models.Country)
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(country.String())
			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				errCon := connect(connectionStatusComponent, loader, &country)
				if errCon != nil {
					errorDialog.Show(errCon)
				}
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
		connectionStatusComponent.Container, // Top
		nil,                                 // Bottom
		nil,                                 // Left
		nil,                                 // Right
		serviceListComponent,                // Rest
	)

	myWindow.SetContent(mainComp)

	myWindow.ShowAndRun()
}
