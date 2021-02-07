# Go Roman Numerals

Convert to and from roman numerals in Go.

```go
package main 

import (
	"fmt"
	"github.com/brandenc40/romannumeral"
)

func ExampleToInt() {
	integer, err := romannumeral.ToInt("IV")
	if err != nil {
		panic(err)
	}
	fmt.Println(integer == 4)
}

func ExampleFromInt() {
	roman, err := romannumeral.FromInt(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(roman == "IV")
}
```
