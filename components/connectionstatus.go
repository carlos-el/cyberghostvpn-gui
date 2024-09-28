package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/carlos-el/cyberghostvpn-gui/models"
)

type ConnectionStatus struct {
	Container *fyne.Container
	Button    *widget.Button
	Label     *widget.Label
}

func NewConnectionStatus(disconnect func()) *ConnectionStatus {
	statusButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		// Defined later
	})
	statusButton.Disable()
	statusLabel := widget.NewLabel("Disconnected")
	statusComponent := container.NewHBox(
		statusButton,
		statusLabel,
	)
	statusButton.OnTapped = func() {
		statusButton.Disable()
		statusLabel.SetText("Disconnected")
		disconnect()
	}

	return &ConnectionStatus{
		Container: statusComponent,
		Button:    statusButton,
		Label:     statusLabel,
	}
}

func (cs *ConnectionStatus) SetConnected(c *models.Country, server string) {
	cs.Button.Enable()

	if server == "" {
		cs.Label.SetText("Disconnected")
	} else if c == nil {
		cs.Label.SetText("Connected to: " + server)
	} else {
		cs.Label.SetText("Connected to: " + server + " - " + c.String())
	}
}
