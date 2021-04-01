// Converts between integers and Roman Numeral strings.
//
// Currently only supports Roman Numerals without viniculum (1-3999) and will throw an error for
// numbers outside of that range. See here for details on viniculum:
// https://en.wikipedia.org/wiki/Roman_numerals#Large_numbers
package romannumeral

import (
	"errors"
	"strings"
)

const (
	_maxRoman = 3999
	_minRoman = 1
)

var (
	InvalidRomanNumeral = errors.New("invalid roman numeral")
	IntegerOutOfBounds  = errors.New("integer must be between 1 and 3999")
)

// numeral describes the value and symbol of a single roman numeral
type numeral struct {
	val int
	sym string
}

// _numerals are all unique numerals ordered from largest to smallest
var _numerals = []numeral{
	{1000, "M"}, {900, "CM"}, {500, "D"},
	{400, "CD"}, {100, "C"}, {90, "XC"},
	{50, "L"}, {40, "XL"}, {10, "X"},
	{9, "IX"}, {5, "V"}, {4, "IV"},
	{1, "I"},
}

// FromInt converts an integer value to a roman numeral string. An error is
// returned if the integer is not between 1 and 3999.
func FromInt(input int) (string, error) {
	if outOfBounds(input) {
		return "", IntegerOutOfBounds
	}
	return intToRoman(input), nil
}

// outOfBounds checks to ensure an input value is valid for roman numerals without the need of
// vinculum (used for values of 4,000 and greater)
func outOfBounds(input int) bool {
	return input < _minRoman || input > _maxRoman
}

func intToRoman(input int) string {
	var output string
	for _, rom := range _numerals {
		for input >= rom.val {
			output += rom.sym
			input -= rom.val
		}
	}
	return output
}

// ToInt converts a roman numeral string to an integer. Roman numerals for numbers
// outside of the range 1 to 3,999 will return an error.
func ToInt(input string) (int, error) {
	if input == "" {
		return 0, nil
	}
	if output, ok := romanToInt(input); ok {
		return output, nil
	}
	return 0, InvalidRomanNumeral
}

func romanToInt(input string) (int, bool) {
	var output int
	for _, rom := range _numerals {
		for strings.HasPrefix(input, rom.sym) {
			output += rom.val
			input = input[len(rom.sym):]
		}
	}
	// If we are still left with input string values then the
	// input was invalid and the bool is returned as False
	return output, len(input) == 0
}
