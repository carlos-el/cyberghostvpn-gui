package commander

import (
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strings"

	"github.com/carlos-el/cyberghostvpn-gui/models"
)

func GetCountryList(srvOpt ServiceType) ([]models.Country, error) {
	cmd := exec.Command("cyberghostvpn", "--"+srvOpt.CommandArg(), "--country-code")
	log.Print(cmd)
	out, err := cmd.Output()
	if err != nil {
		return []models.Country{}, &ErrCommandSysExecution{Msg: "in commander GetCountryList, could not run command", Err: err}
	}
	log.Print(string(out))

	countryList, parseErr := countryListParser(string(out))
	if parseErr != nil {
		return []models.Country{}, fmt.Errorf("in commander GetCountryList: %w", parseErr)
	}

	sort.Slice(countryList, func(i, j int) bool {
		return countryList[i].Name < countryList[j].Name
	})

	return countryList, nil
}

func Connect(c *models.Country, vpnOpt VpnProtocol, transOpt TransmissionProtocol, srvOpt ServiceType) (string, error) {
	command := fmt.Sprintf(
		"cyberghostvpn --%s --country-code %s --%s --%s --connect",
		srvOpt.CommandArg(),
		c.Code,
		vpnOpt.CommandArg(),
		transOpt.CommandArg(),
	)
	log.Print(command)
	cmd := exec.Command("/bin/sh", "-c", command)
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		return "", &ErrCommandSysExecution{Msg: "in commander Connect, could not run command: " + command, Err: cmdErr}
	}

	sudoErr := DetectErrSudoRequiredInMsg(string(out))
	if sudoErr != nil {
		return "", sudoErr
	}

	log.Print(string(out))
	server, parseErr := parseServerFromConnectMsg(string(out))
	if parseErr != nil {
		return "", fmt.Errorf("in commander func Connect: %w", parseErr)
	}

	return server, nil
}

func Disconnect() error {
	command := "cyberghostvpn --stop"
	cmd := exec.Command("/bin/sh", "-c", command)
	log.Print(command)
	out, err := cmd.Output()
	if err != nil {
		return &ErrCommandSysExecution{Msg: "in commander Disconnect, could not run command", Err: err}
	}
	sudoErr := DetectErrSudoRequiredInMsg(string(out))
	if sudoErr != nil {
		return sudoErr
	}
	log.Print(string(out))
	return nil
}

func CheckConnection() (bool, error) {
	cmd := exec.Command("cyberghostvpn", "--status")
	log.Print(cmd)
	out, err := cmd.Output()
	if err != nil {
		return false, &ErrCommandSysExecution{Msg: "in commander CheckConnection, could not run command", Err: err}
	}

	outText := string(out)
	log.Print(outText)

	if strings.HasPrefix(outText, "No VPN") {
		return false, nil
	}
	return true, nil
}
