package main

import (
	"fmt"

	"LearnGolang/src/gopl/ch6/intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(9)
	x.Add(144)
	fmt.Println(x.String())
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())
	x.Union(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))
}
