package commander

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/carlos-el/cyberghostvpn-gui/models"
)

const ServerNameRegex = "^([a-z]{2,50})-s([0-9]{3,4})-i([0-9]{2,3})$"

func countryListParser(stringTable string) ([]models.Country, error) {
	list := make([]models.Country, 0)
	// Split the table by line jumps
	stringList := strings.Split(stringTable, "\n")
	// If the command comes with an additional line at the beginning for auth the start parsing one line after it
	dataStartLine := 3
	if strings.Contains(stringList[0], "Perform") {
		dataStartLine = 4
	}
	// For each row with countries get them correctly by splitting and trimming the chars and add them to the list
	for i := dataStartLine; i < len(stringList)-2; i++ {
		aux := strings.Split(stringList[i], "|")
		countryNumber, err := strconv.Atoi(strings.TrimSpace(aux[1]))
		if err != nil {
			return []models.Country{}, &ErrCommandParse{
				Msg:  "in commander countryListParser, could not parse country code",
				Text: stringTable,
				Err:  err,
			}
		}

		country := models.Country{
			Number: countryNumber,
			Name:   strings.TrimSpace(aux[2]),
			Code:   strings.TrimSpace(aux[3]),
		}
		list = append(list, country)
	}

	return list, nil
}

func parseServerFromConnectMsg(msg string) (string, error) {
	stringList := strings.Split(msg, "\n")
	secondLine := stringList[1]
	server := strings.TrimSpace(secondLine[17:])

	matched, err := regexp.MatchString(ServerNameRegex, server)
	if !matched || err != nil {
		return "", &ErrCommandParse{
			Msg:  "in commander parseServerFromConnectMsg, server name could not be parsed",
			Text: msg,
			Err:  err,
		}
	}

	return server, nil
}
