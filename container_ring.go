package main

import (
	"container/ring"
	"fmt"
	"math"
)

func main() {

	// creating a new ring with 4 element
	myRing := ring.New(4)

	// adding element to the ring

	for i := 1; i <= 4; i++ {
		myRing.Value = math.Pow(float64(i), float64(i))
		myRing = myRing.Next()
	}

	// itering over the ring and printing its contents
	fmt.Println("Element of the ring:")
	myRing.Do(func(x interface{}) {
		fmt.Println(x)
	})

}
