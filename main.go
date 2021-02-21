package awesomeProject

import "strings"

func ConvertToRoman(number int) string {
	var result strings.Builder

	for number > 0 {
		switch {
		case number > 4:
			result.WriteString("V")
			number -= 5
		case number > 3:
			result.WriteString("IV")
			number -= 4
		default:
			result.WriteString("I")
			number--
		}
	}

	return result.String()
}
