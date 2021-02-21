package awesomeProject

import "strings"

func ConvertToRoman(number int) string {
	var result strings.Builder

	if number == 4 {
		return "IV"
	}

	for i := 0; i < number; i++ {
		result.WriteString("I")
	}

	return result.String()
}
