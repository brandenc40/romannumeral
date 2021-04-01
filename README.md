# Go Roman Numerals

[![Go Reference](https://pkg.go.dev/badge/github.com/brandenc40/romannumeral.svg)](https://pkg.go.dev/github.com/brandenc40/romannumeral)
[![codecov](https://codecov.io/gh/brandenc40/romannumeral/branch/master/graph/badge.svg?token=AS7IBSTE36)](https://codecov.io/gh/brandenc40/romannumeral)
    
Quickly and efficiently convert to and from roman numerals in Go.

### Benchmark Results

```sh
goos: darwin
goarch: arm64
pkg: github.com/brandenc40/romannumeral
BenchmarkToRomanNumeral
BenchmarkToRomanNumeral-8   	10997846	       108.5 ns/op
BenchmarkToInteger
BenchmarkToInteger-8        	18861016	        62.29 ns/op
```

### Example

```go
package main

import (
	"fmt"
	rom "github.com/brandenc40/romannumeral"
)

func ExampleRomanToInt() {
	integer, err := rom.ToInt("IV")
	if err != nil {
		panic(err)
	}
	fmt.Println(integer == 4)
}

func ExampleIntToRoman() {
	roman, err := rom.FromInt(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(roman == "IV")
}
```
