package adress

import "strings"

// AdressType CHECK ADRESS TYPE AND RETURN IT
func AdressType(adress string) string {
	validTypes := []string{"street", "avenue", "road", "highway"}
	formatAdress := strings.ToLower(adress)
	firstElement := strings.Split(formatAdress, " ")[0]

	validAdress := false
	for _, tipo := range validTypes {
		if tipo == firstElement {
			validAdress = true
		}
	}

	if validAdress {
		return strings.Title(firstElement)
	}

	return "Invalid Adress"
}
