package main

import (
	"fmt"
	"sort"
)

func main() {

	// sorting slice of integer
	intSlice := []int{2, 4, 3, 6, 3, 1}
	sort.Ints(intSlice)
	fmt.Println("Sorted slice of integer:", intSlice)

	// sorting slice of strings
	strSlice := []string{"banana", "grape", "apple", "pear"}
	sort.Strings(strSlice)
	fmt.Println("Sorted slice of string:", strSlice)
}
