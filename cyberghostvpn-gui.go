package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

func connect(
	status *components.ConnectionStatus,
	loader *dialog.CustomDialog,
	c *models.Country,
	vpnOptionFunc func() commander.VpnProtocol,
	transOption func() commander.TransmissionProtocol,
	srvOption func() commander.ServiceType,
) error {
	loader.Show()
	defer loader.Hide()

	server, err := commander.Connect(c, vpnOptionFunc(), transOption(), srvOption())
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}

	status.SetConnected(c, server)
	return nil
}

func GetServerList(serviceType commander.ServiceType, errorDialog *components.ErrorDialog, loader *dialog.CustomDialog) []models.Country {
	loader.Show()
	defer loader.Hide()
	var countries, err = commander.GetCountryList(serviceType)
	if err != nil {
		errorDialog.Show(err)
	}
	return countries
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

	// Connection options selectors
	connectionOptionsComponent := components.NewConnectionOptions()

	// Service type selector
	serviceTypeOptions := []string{
		commander.Traffic.String(),
		commander.Torrent.String(),
		commander.Streaming.String(),
	}
	inputServiceType := widget.NewSelect(serviceTypeOptions, nil)
	inputServiceType.SetSelected(serviceTypeOptions[0])

	// Server list array
	serverList := GetServerList(commander.ServiceType(inputServiceType.SelectedIndex()+1), errorDialog, loader)
	// Service list component
	list := widget.NewList(
		func() int {
			return len(serverList)
		},
		func() fyne.CanvasObject {
			b := widget.NewButtonWithIcon("", theme.LoginIcon(), func() {
				// Defined later
			})
			return container.NewHBox(widget.NewLabel("template"), layout.NewSpacer(), b)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			country := serverList[i]
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(country.String())
			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				errCon := connect(
					connectionStatusComponent,
					loader,
					&country,
					connectionOptionsComponent.GetVpnOption,
					connectionOptionsComponent.GetTransmissionOption,
					func() commander.ServiceType {
						return commander.ServiceType(inputServiceType.SelectedIndex() + 1)
					},
				)
				if errCon != nil {
					errorDialog.Show(errCon)
				}
			}
		},
	)

	// Update service list on service type change
	inputServiceType.OnChanged = func(s string) {
		log.Print("Changed selected service type: ", s)
		serverList = GetServerList(commander.ServiceType(inputServiceType.SelectedIndex()+1), errorDialog, loader)
		list.Refresh()
	}

	// List component
	listTitle := widget.NewRichTextFromMarkdown("### Select a server to connect to: ")
	serviceListComponent := container.NewBorder(
		container.NewVBox(listTitle, inputServiceType), // Top
		nil,  // Bottom
		nil,  // Left
		nil,  // Right
		list, // Rest
	)

	// Main component
	mainComp := container.NewBorder(
		container.NewVBox(
			connectionStatusComponent.Container,
			canvas.NewText("", color.White),
			connectionOptionsComponent.Container,
			canvas.NewText("", color.White),
		), // Top
		nil,                  // Bottom
		nil,                  // Left
		nil,                  // Right
		serviceListComponent, // Rest
	)

	myWindow.SetContent(mainComp)

	myWindow.ShowAndRun()
}
