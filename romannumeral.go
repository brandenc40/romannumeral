// Converts between integers and Roman Numeral strings.
//
// Currently only supports Roman Numerals without viniculum (1-3999) and will throw an error for
// numbers outside of that range. See here for details on viniculum:
// https://en.wikipedia.org/wiki/Roman_numerals#Large_numbers
package romannumeral

import (
	"bytes"
	"errors"
)

// numeral describes the value and symbol of a single roman numeral
type numeral struct {
	val int
	sym []byte
}

var (
	// InvalidRomanNumeral - error for when a roman numeral string provided is not a valid roman numeral
	InvalidRomanNumeral = errors.New("invalid roman numeral")
	// IntegerOutOfBounds - error for when the integer provided is invalid and unable to be converted to a roman numeral
	IntegerOutOfBounds = errors.New("integer must be between 1 and 3999")

	// all unique numerals ordered from largest to smallest
	nums = []numeral{
		{1000, []byte("M")},
		{900, []byte("CM")},
		{500, []byte("D")},
		{400, []byte("CD")},
		{100, []byte("C")},
		{90, []byte("XC")},
		{50, []byte("L")},
		{40, []byte("XL")},
		{10, []byte("X")},
		{9, []byte("IX")},
		{5, []byte("V")},
		{4, []byte("IV")},
		{1, []byte("I")},
	}

	// lookup arrays used for converting from an int to a roman numeral extremely quickly.
	// method from here: https://rosettacode.org/wiki/Roman_numerals/Encode#Go
	r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	r3 = []string{"", "M", "MM", "MMM"}
)

// IntToString converts an integer value to a roman numeral string. An error is
// returned if the integer is not between 1 and 3999.
func IntToString(input int) (string, error) {
	if outOfBounds(input) {
		return "", IntegerOutOfBounds
	}
	return intToRoman(input), nil
}

// IntToBytes converts an integer value to a roman numeral byte array. An error is
// returned if the integer is not between 1 and 3999.
func IntToBytes(input int) ([]byte, error) {
	str, err := IntToString(input)
	return []byte(str), err
}

// outOfBounds checks to ensure an input value is valid for roman numerals without the need of
// vinculum (used for values of 4,000 and greater)
func outOfBounds(input int) bool {
	return input < 1 || input > 3999
}

func intToRoman(n int) string {
	// This is efficient in Go. The 4 operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return r3[n%1e4/1e3] + r2[n%1e3/1e2] + r1[n%100/10] + r0[n%10]
}

// StringToInt converts a roman numeral string to an integer. Roman numerals for numbers
// outside of the range 1 to 3,999 will return an error. Empty strings will return 0
// with no error thrown.
func StringToInt(input string) (int, error) {
	return BytesToInt([]byte(input))
}

// BytesToInt converts a roman numeral byte array to an integer. Roman numerals for numbers
// outside of the range 1 to 3,999 will return an error. Nil or empty []byte will return 0
// with no error thrown.
func BytesToInt(input []byte) (int, error) {
	if input == nil || len(input) == 0 {
		return 0, nil
	}
	if output, ok := romanToInt(input); ok {
		return output, nil
	}
	return 0, InvalidRomanNumeral
}

func romanToInt(input []byte) (int, bool) {
	var output int
	for _, n := range nums {
		for bytes.HasPrefix(input, n.sym) {
			output += n.val
			input = input[len(n.sym):]
		}
	}
	// if we are still left with input string values then the
	// input was invalid and the bool is returned as false
	return output, len(input) == 0
}
