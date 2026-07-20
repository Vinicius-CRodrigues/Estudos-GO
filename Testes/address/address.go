package address

import (
	"strings"
)

func TypeOfAddresses(address string) string {
	addresses := []string{"Rua", "Quadra", "Chacara", "Avenida"}

	firstWordOfAddress := strings.Split(address, " ")[0]

	valid := false

	for _, word := range addresses {
		if word == firstWordOfAddress {
			valid = true
		}

	}

	if valid == true {

		return firstWordOfAddress
	}

	return "Invalid type"
}
