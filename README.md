# Go Roman Numerals

[![Go Reference](https://pkg.go.dev/badge/github.com/brandenc40/romannumeral.svg)](https://pkg.go.dev/github.com/brandenc40/romannumeral)
[![codecov](https://codecov.io/gh/brandenc40/romannumeral/branch/master/graph/badge.svg?token=AS7IBSTE36)](https://codecov.io/gh/brandenc40/romannumeral)
    
## Quickly and efficiently convert to and from roman numerals in Go.

A reliable module using the most efficient methods possible for converting between
roman numerals and integers in Go. Algorithms adopted from [here](https://rosettacode.org/wiki/Roman_numerals).

### Benchmark Results

```sh
goos: darwin
goarch: arm64
pkg: github.com/brandenc40/romannumeral
BenchmarkIntToString-8          56474846                20.84 ns/op            0 B/op          0 allocs/op
BenchmarkIntToBytes-8           48157634                24.36 ns/op            0 B/op          0 allocs/op
BenchmarkStringToInt-8          17584252                67.28 ns/op            0 B/op          0 allocs/op
BenchmarkBytesToInt-8           18343551                64.77 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/brandenc40/romannumeral      6.111s
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
