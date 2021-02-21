package awesomeProject

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(number int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for number >= numeral.Value {
			result.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}

	return result.String()
}
