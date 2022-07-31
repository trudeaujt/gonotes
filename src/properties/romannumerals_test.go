package properties

import (
	"fmt"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		Arabic uint16
		Roman  string
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
)

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			want := test.Arabic
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

/*
There have been a few rules in the domain of Roman Numerals that we have worked with in this chapter
Can't have more than 3 consecutive symbols
Only I (1), X (10) and C (100) can be "subtractors"
Taking the result of ConvertToRoman(N) and passing it to ConvertToArabic should return us N

So far, the tests we have written have been 'example' based tests where we provide the tooling some examples around our code to verify.

What if we could take these rules that we know about our domain and somehow exercise them against our code?

Property based tests help us do this by throwing random data at our code and verifying that the rules we describe always hold true.
*/
func TestPropertiesOfConversion(t *testing.T) {
	//This will check that we get the number we originally had when converting then converting back.
	//However, using int will use very large (and negative) values that our implementation cannot handle.
	//So we can use uint16 instead.
	assertion := func(arabic uint16) bool {
		//This is the largest value our implementation can support.
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	//quick.Check will run against a number of random inputs.
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
