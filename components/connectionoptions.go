package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/carlos-el/cyberghostvpn-gui/commander"
)

type ConnectionOptionsSelector struct {
	Container                 *fyne.Container
	InputVpnProtocol          *widget.Select
	InputTransmissionProtocol *widget.Select
}

func NewConnectionOptions() *ConnectionOptionsSelector {
	connectionOptionsHeaderText := widget.NewRichTextFromMarkdown("### Connection options: ")
	vpnProtocolOptions := []string{commander.OpenVpn.String(), commander.WireGuard.String()}
	inputVpnProtocol := widget.NewSelect(vpnProtocolOptions, nil)
	inputVpnProtocol.SetSelected(vpnProtocolOptions[0])
	transmissionProtocolOptions := []string{commander.Tcp.String(), commander.Udp.String()}
	inputTransmissionProtocol := widget.NewSelect(transmissionProtocolOptions, nil)
	inputTransmissionProtocol.SetSelected(transmissionProtocolOptions[0])

	connectionOptionsForm := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("VPN protocol"),
		inputVpnProtocol,
		widget.NewLabel("Transmission protocol"),
		inputTransmissionProtocol,
	)
	connectionOptionsComponent := container.NewVBox(
		connectionOptionsHeaderText,
		connectionOptionsForm,
	)

	return &ConnectionOptionsSelector{
		Container:                 connectionOptionsComponent,
		InputVpnProtocol:          inputVpnProtocol,
		InputTransmissionProtocol: inputTransmissionProtocol,
	}
}

func (co *ConnectionOptionsSelector) GetVpnOption() commander.VpnProtocol {
	return commander.VpnProtocol(co.InputVpnProtocol.SelectedIndex() + 1)
}
func (co *ConnectionOptionsSelector) GetTransmissionOption() commander.TransmissionProtocol {
	return commander.TransmissionProtocol(co.InputTransmissionProtocol.SelectedIndex() + 1)
}
