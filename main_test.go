package awesomeProject

import "testing"

func assertRomanNumeral(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConvertToRoman(t *testing.T) {
	cases := []struct {
		Description string
		Number      int
		RomanNumber string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Number)
			assertRomanNumeral(t, got, test.RomanNumber)
		})
	}
}
