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
	sym []byte
}

// _numerals are all unique numerals ordered from largest to smallest
var _numerals = []numeral{
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

// IntToString converts an integer value to a roman numeral string. An error is
// returned if the integer is not between 1 and 3999.
func IntToString(input int) (string, error) {
	if outOfBounds(input) {
		return "", IntegerOutOfBounds
	}
	return string(intToRoman(input)), nil
}

// outOfBounds checks to ensure an input value is valid for roman numerals without the need of
// vinculum (used for values of 4,000 and greater)
func outOfBounds(input int) bool {
	return input < _minRoman || input > _maxRoman
}

func intToRoman(input int) []byte {
	var output []byte
	for _, rom := range _numerals {
		for input >= rom.val {
			output = append(output, rom.sym...)
			input -= rom.val
		}
	}
	return output
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
	for _, rom := range _numerals {
		for bytes.HasPrefix(input, rom.sym) {
			output += rom.val
			input = input[len(rom.sym):]
		}
	}
	// If we are still left with input string values then the
	// input was invalid and the bool is returned as False
	return output, len(input) == 0
}
