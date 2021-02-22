package awesomeProject

import (
	"fmt"
	"testing"
)

func assertRomanNumeral(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertArabicNumeral(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

var cases = []struct {
	Number      int
	RomanNumber string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d should be %q", test.Number, test.RomanNumber), func(t *testing.T) {
			got := ConvertToRoman(test.Number)
			assertRomanNumeral(t, got, test.RomanNumber)
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases[:1] {
		t.Run(fmt.Sprintf("%q should be %d", test.RomanNumber, test.Number), func(t *testing.T) {
			got := ConvertToArabic(test.RomanNumber)
			assertArabicNumeral(t, got, test.Number)
		})
	}
}
