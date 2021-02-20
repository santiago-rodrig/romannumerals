package awesomeProject

import "strings"

func ConvertToRoman(number int) string {
	var result strings.Builder

	for i := 0; i < number; i++ {
		result.WriteString("I")
	}

	return result.String()
}
