package romannumeral

import (
	"errors"
	"strings"
)

type numeral struct {
	val int
	sym string
}

var (
	InvalidRomanNumeral = errors.New("invalid roman numeral")
	IntegerOutOfBounds  = errors.New("integer must be between 1 and 3999")
)

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

func outOfBounds(input int) bool {
	return input < 1 || input > 3999
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
// outside of the range 1 to 3999 will return an error.
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
