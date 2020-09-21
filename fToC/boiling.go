package main

import (
	"fmt"
)

var i, j, k int
var b, f, s = true, 2.3, "four"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	// in, err := os.Open("infile")
	// out, err := os.Create("outfile")
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
