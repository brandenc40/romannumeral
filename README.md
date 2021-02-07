# Go Roman Numerals

Convert to and from roman numerals in Go.

```go
func ExampleToInt() {
	integer, err := ToInt("IV")
	if err != nil {
		panic(err)
	}
	fmt.Println(integer == 4)
}

func ExampleFromInt() {
	roman, err := FromInt(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(roman == "IV")
}
```
