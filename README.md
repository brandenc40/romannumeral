# Go Roman Numerals

[![Go Reference](https://pkg.go.dev/badge/github.com/brandenc40/romannumeral.svg)](https://pkg.go.dev/github.com/brandenc40/romannumeral)
[![codecov](https://codecov.io/gh/brandenc40/romannumeral/branch/master/graph/badge.svg?token=AS7IBSTE36)](https://codecov.io/gh/brandenc40/romannumeral)
    
Quickly and efficiently convert to and from roman numerals in Go.

### Benchmark Results

```sh
goos: darwin
goarch: arm64
pkg: github.com/brandenc40/romannumeral
BenchmarkIntToString-8          18174510                65.95 ns/op           24 B/op          2 allocs/op
BenchmarkStringToInt-8          17507576                67.26 ns/op            0 B/op          0 allocs/op
BenchmarkBytesToInt-8           18342873                64.75 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/brandenc40/romannumeral      5.151s
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
