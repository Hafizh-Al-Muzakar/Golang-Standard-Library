package main

import (
	"fmt"
	"math"
)

func main() {

	// calculate square root

	x := 9.0
	sqrt := math.Sqrt(x)

	fmt.Println("Square Root of", x, "is", sqrt)

	// calculate power
	base := 2.0
	exp := 3.0
	pow := math.Pow(base, exp)
	fmt.Println("Power of", base, "raised to", exp, "is", pow)

	// calculate absolute value
	y := -4.0
	abs := math.Abs(y)
	fmt.Println("Absolute value of", y, "is", abs)
}
