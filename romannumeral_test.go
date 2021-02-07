package romannumeral

import (
	"testing"
)

var testCases = map[string]int{
	"I": 1, "II": 2,
	"III": 3, "IV": 4,
	"V": 5, "VI": 6,
	"VII": 7, "VIII": 8,
	"IX": 9, "X": 10,
	"XIV": 14, "XIX": 19,
	"XX": 20, "XXXIII": 33,
	"XXXIV": 34, "XXXIX": 39,
	"XLIX": 49, "L": 50,
	"LXXXIX": 89, "XCIX": 99,
	"CXLIX": 149, "CCCXLIX": 349,
	"CDLVI": 456, "D": 500,
	"DCIV": 604, "DCCLXXXIX": 789,
	"DCCCXLIX": 849, "CMIV": 904,
	"MVII": 1007, "MLXVI": 1066,
	"MDCCLXXVI": 1776, "MMDCCCVI": 2806,
	"MMCMXCIX": 2999, "MMXXI": 2021,
	"MMMCMLXXIX": 3979, "MMMCMXCIX": 3999,
}

func TestFromInt(t *testing.T) {
	for expected, input := range testCases {
		out, err := FromInt(input)
		if err != nil {
			t.Errorf("FromInt(%d) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("FromInt(%d) = %s; want %s", input, out, expected)
		}
	}
	_, err := FromInt(100000)
	if err == nil {
		t.Errorf("FromInt(%d) expected an error", 100000)
	}
	_, err = FromInt(0)
	if err == nil {
		t.Errorf("FromInt(%d) expected an error", 0)
	}
}

func TestToInteger(t *testing.T) {
	for input, expected := range testCases {
		out, err := ToInt(input)
		if err != nil {
			t.Errorf("ToInt(%s) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("ToInt(%s) = %d; want %d", input, out, expected)
		}
	}
	_, err := ToInt("IVCMXCIX")
	if err == nil {
		t.Errorf("ToInt(%d) expected an error", 0)
	}
}

func BenchmarkToRomanNumeral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = FromInt(2999)
	}
}

func BenchmarkToInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ToInt("DCCCXLIX")
	}
}
