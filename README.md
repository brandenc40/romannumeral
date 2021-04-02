# Go Roman Numerals

[![Go Reference](https://pkg.go.dev/badge/github.com/brandenc40/romannumeral.svg)](https://pkg.go.dev/github.com/brandenc40/romannumeral)
[![codecov](https://codecov.io/gh/brandenc40/romannumeral/branch/master/graph/badge.svg?token=AS7IBSTE36)](https://codecov.io/gh/brandenc40/romannumeral)
    
Quickly and efficiently convert to and from roman numerals in Go.

### Benchmark Results

```sh
goos: darwin
goarch: arm64
pkg: github.com/brandenc40/romannumeral
BenchmarkIntToString-8          18160285                66.43 ns/op           24 B/op          2 allocs/op
BenchmarkIntToBytes-8           18630670                63.68 ns/op           24 B/op          2 allocs/op
BenchmarkStringToInt-8          17537406                67.26 ns/op            0 B/op          0 allocs/op
BenchmarkBytesToInt-8           18248730                64.83 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/brandenc40/romannumeral      6.283s

```

### Example

```go
package main

import (
	"fmt"
	rom "github.com/brandenc40/romannumeral"
)

func ExampleStringToInt() {
	integer, err := rom.StringToInt("IV")
	if err != nil {
		panic(err)
	}
	fmt.Println(integer == 4) // True
}

func ExampleBytesToInt() {
	integer, err := rom.BytesToInt([]byte("IV"))
	if err != nil {
		panic(err)
	}
	fmt.Println(integer == 4) // True
}

func ExampleIntToString() {
	roman, err := rom.IntToString(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(roman == "IV") // True
}

func ExampleIntToBytes() {
	roman, err := rom.IntToBytes(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(roman) == "IV") // True
}
```
