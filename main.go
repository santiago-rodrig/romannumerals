package awesomeProject

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)

	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
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

func couldBeSubtractive(index int, currentSymbol byte, roman string) bool {
	return index+1 < len(roman) && currentSymbol == 'I'
}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]
			potentialNumber := []byte{symbol, nextSymbol}

			if value := allRomanNumerals.ValueOf(potentialNumber...); value != 0 {
				total += value
				i++
			} else {
				total++
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}

	return total
}
