package commander

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/carlos-el/cyberghostvpn-gui/models"
)

func countryListParser(stringTable string) []models.Country {
	list := make([]models.Country, 0)
	// Split the table by line jumps
	stringList := strings.Split(stringTable, "\n")
	// For each row with countries get them correctly by splitting and trimming the chars and add them to the list
	for i := 3; i < len(stringList)-2; i++ {
		aux := strings.Split(stringList[i], "|")
		countryNumber, _ := strconv.Atoi(strings.TrimSpace(aux[1]))

		country := models.Country{
			Number: countryNumber,
			Name:   strings.TrimSpace(aux[2]),
			Code:   strings.TrimSpace(aux[3]),
		}
		list = append(list, country)
	}

	return list
}

func GetCountryList() ([]models.Country, error) {
	cmd := exec.Command("cyberghostvpn", "--traffic", "--country-code")
	out, err := cmd.Output()
	if err != nil {
		return make([]models.Country, 0), fmt.Errorf("could not run command: %w", err)
	}

	countryList := countryListParser(string(out))

	return countryList, nil
}
