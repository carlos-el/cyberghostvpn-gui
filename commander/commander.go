package commander

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/carlos-el/cyberghostvpn-gui/models"
)

func GetCountryList() ([]models.Country, error) {
	cmd := exec.Command("cyberghostvpn", "--traffic", "--country-code")
	out, err := cmd.Output()
	if err != nil {
		return []models.Country{}, &ErrCommandSysExecution{Msg: "in commander GetCountryList, could not run command", Err: err}
	}
	log.Print(string(out))

	countryList, parseErr := countryListParser(string(out))
	if parseErr != nil {
		return []models.Country{}, fmt.Errorf("in commander GetCountryList: %w", parseErr)
	}

	return countryList, nil
}

func Connect(c *models.Country) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", "cyberghostvpn --traffic --country-code "+c.Code+" --connect")
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		return "", &ErrCommandSysExecution{Msg: "in commander Connect, could not run command", Err: cmdErr}
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
	cmd := exec.Command("/bin/sh", "-c", "cyberghostvpn --stop")
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
