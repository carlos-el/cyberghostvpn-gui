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

func disconnect(statusButton *widget.Button, statusLabel *widget.Label, loader *dialog.CustomDialog) error {
	loader.Show()
	defer loader.Hide()

	err := commander.Disconnect()
	if err != nil {
		return fmt.Errorf("could not disconnect: %w", err)
	}

	statusButton.Disable()
	setConnectionLabel(statusLabel, nil, "")
	return nil
}

func connect(statusButton *widget.Button, statusLabel *widget.Label, loader *dialog.CustomDialog, c *models.Country) error {
	loader.Show()
	defer loader.Hide()

	server, err := commander.Connect(c)
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}

	statusButton.Enable()
	setConnectionLabel(statusLabel, c, server)
	return nil
}

func showErrorDialog(errorDialog *dialog.CustomDialog, errorDialogText *widget.Label, err error) {
	errorDialogText.SetText(err.Error())
	errorDialog.Show()
}

func main() {
	a := app.New()
	myWindow := a.NewWindow("CyberGhost GUI")
	myWindow.Resize(fyne.NewSize(400, 600))

	// Error dialog
	errorDialogIntro := widget.NewLabel("An unexpected error ocurred, \nthe application will be closing.")
	errorDialogText := widget.NewLabel("Error message")
	errorDialogTextContainer := container.NewScroll(errorDialogText)
	errorDialogTextContainer.SetMinSize(fyne.NewSize(100, 200))
	errorDialogButton := widget.NewButton("Close", func() {
		a.Quit()
	})
	errorDialogContent := container.NewVBox(errorDialogIntro, errorDialogTextContainer, errorDialogButton)
	errorDialog := dialog.NewCustomWithoutButtons("An error ocurred:", errorDialogContent, myWindow)
	// Loader dialog to block interaction
	loader := dialog.NewCustomWithoutButtons("", widget.NewLabel("Loading..."), myWindow)
	// Status component
	statusButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		// Defined later
	})
	statusButton.Disable()
	statusLabel := widget.NewLabel("Disconnected")
	statusButton.OnTapped = func() {
		err := disconnect(statusButton, statusLabel, loader)
		if err != nil {
			showErrorDialog(errorDialog, errorDialogText, err)
		}
	}
	statusComponent := container.NewHBox(
		statusButton,
		statusLabel,
	)

	// Get data
	var countries, err = commander.GetCountryList()
	if err != nil {
		showErrorDialog(errorDialog, errorDialogText, err)
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
				showErrorDialog(errorDialog, errorDialogText, errCasting)
			}
			country := c.(models.Country)
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(country.String())
			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				errCon := connect(statusButton, statusLabel, loader, &country)
				if errCon != nil {
					showErrorDialog(errorDialog, errorDialogText, errCon)
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
		statusComponent,      // Top
		nil,                  // Bottom
		nil,                  // Left
		nil,                  // Right
		serviceListComponent, // Rest
	)

	myWindow.SetContent(mainComp)

	myWindow.ShowAndRun()
}
