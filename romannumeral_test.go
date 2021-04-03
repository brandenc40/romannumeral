package romannumeral

import (
	"fmt"
	"testing"
)

var testCases = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12,
	"XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17,
	"XVIII": 18, "XIX": 19, "XX": 20, "XXXI": 31, "XXXII": 32,
	"XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37,
	"XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42,
	"XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47,
	"XLVIII": 48, "XLIX": 49, "L": 50, "LXXXIX": 89, "XC": 90,
	"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95,
	"XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	"CI": 101, "CII": 102, "CIII": 103, "CIV": 104, "CV": 105,
	"CVI": 106, "CVII": 107, "CVIII": 108, "CIX": 109, "CXLIX": 149,
	"CCCXLIX": 349, "CDLVI": 456, "D": 500, "DCIV": 604, "DCCLXXXIX": 789,
	"DCCCXLIX": 849, "CMIV": 904, "M": 1000, "MVII": 1007, "MLXVI": 1066,
	"MCCXXXIV": 1234, "MDCCLXXVI": 1776, "MMXXI": 2021, "MMDCCCVI": 2806,
	"MMCMXCIX": 2999, "MMM": 3000, "MMMCMLXXIX": 3979, "MMMCMXCIX": 3999,
}

func TestIntToString(t *testing.T) {
	for expected, input := range testCases {
		out, err := IntToString(input)
		if err != nil {
			t.Errorf("IntToString(%d) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("IntToString(%d) = %s; want %s", input, out, expected)
		}
	}
	_, err := IntToString(100000)
	if err == nil {
		t.Errorf("IntToString(%d) expected an error", 100000)
	}
	_, err = IntToString(0)
	if err == nil {
		t.Errorf("IntToString(%d) expected an error", 0)
	}
}

func TestIntToBytes(t *testing.T) {
	for expected, input := range testCases {
		out, err := IntToBytes(input)
		if err != nil {
			t.Errorf("IntToBytes(%d) returned an error %s", input, err.Error())
		}
		if len(out) != len([]byte(expected)) {
			t.Errorf("len(IntToBytes(%d)) = %d; want %d", input, len(out), len([]byte(expected)))
		}
		for i, char := range out {
			if char != expected[i] {
				t.Errorf("IntToBytes(%d) = %s; want %s", input, string(out), expected)
			}
		}
	}
	_, err := IntToBytes(100000)
	if err == nil {
		t.Errorf("IntToBytes(%d) expected an error", 100000)
	}
	_, err = IntToBytes(0)
	if err == nil {
		t.Errorf("IntToBytes(%d) expected an error", 0)
	}
}

func TestStringToInt(t *testing.T) {
	for input, expected := range testCases {
		out, err := StringToInt(input)
		if err != nil {
			t.Errorf("StringToInt(%s) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("StringToInt(%s) = %d; want %d", input, out, expected)
		}
	}
	_, err := StringToInt("IVCMXCIX")
	if err == nil {
		t.Error("StringToInt(IVCMXCIX) expected an error")
	}

	val, err := StringToInt("")
	if val != 0 {
		t.Errorf("StringToInt(\"\") = %d; want 0", val)
	}
	if err != nil {
		t.Errorf("StringToInt(\"\") returned an error %s", err.Error())
	}
}

func TestBytesToInt(t *testing.T) {
	for input, expected := range testCases {
		out, err := BytesToInt([]byte(input))
		if err != nil {
			t.Errorf("StringToInt(%s) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("StringToInt(%s) = %d; want %d", input, out, expected)
		}
	}
	_, err := BytesToInt([]byte("IVCMXCIX"))
	if err == nil {
		t.Error("BytesToInt(IVCMXCIX) expected an error")
	}

	var in []byte
	val, err := BytesToInt(in)
	if val != 0 {
		t.Errorf("BytesToInt(nil) = %d; want 0", val)
	}
	if err != nil {
		t.Errorf("BytesToInt(nil) returned an error %s", err.Error())
	}

	in = []byte("")
	val, err = BytesToInt(in)
	if val != 0 {
		t.Errorf("BytesToInt([]byte(\"\")) = %d; want 0", val)
	}
	if err != nil {
		t.Errorf("BytesToInt([]byte(\"\")) returned an error %s", err.Error())
	}
}

func BenchmarkIntToString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = IntToString(3999)
	}
}

func BenchmarkIntToBytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = IntToBytes(3999)
	}
}

func BenchmarkStringToInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = StringToInt("MMMCMXCIX")
	}
}

func BenchmarkBytesToInt(b *testing.B) {
	in := []byte("MMMCMXCIX")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = BytesToInt(in)
	}
}

func ExampleStringToInt() {
	integer, err := StringToInt("IV")
	if err != nil {
		panic(err)
	}
	fmt.Println(integer == 4) // True
}

func ExampleBytesToInt() {
	input := []byte("IV")
	integer, err := BytesToInt(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(integer == 4) // True
}

func ExampleIntToString() {
	roman, err := IntToString(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(roman == "IV") // True
}

func ExampleIntToBytes() {
	roman, err := IntToBytes(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(roman) == "IV") // True
}
