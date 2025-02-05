package main

import (
	"fmt"
	"reflect"
)

func printTypeAndValue(v interface{}) {
	//get type of the variable
	t := reflect.TypeOf(v)
	fmt.Println("Type:", t)

	//get the value of the variable
	val := reflect.ValueOf(v)

	//check if the variable is a slice or map
	if t.Kind() == reflect.Slice {
		//iterate over the elements of the slice
		fmt.Println("Value:")
		for i := 0; i < val.Len(); i++ {
			fmt.Println("\t", val.Index(i))
		}
	} else if t.Kind() == reflect.Map {
		fmt.Println("Value (Map):")
		keys := val.MapKeys()
		for key, value := range keys {
			fmt.Println("\t", key, ":", value)
		}
	} else {
		fmt.Println("Value:", val)
	}
}

func main() {

	//define variables of different types

	var (
		num     = 10
		str     = "Hello Mama"
		slice   = []int{1, 2, 3, 4, 5}
		mapping = map[string]int{"a": 1, "b": 2}
	)

	fmt.Println("Variable 'num':")
	printTypeAndValue(num)

	fmt.Println("Variable 'str':")
	printTypeAndValue(str)

	fmt.Println("Variable 'slice':")
	printTypeAndValue(slice)

	fmt.Println("Variable 'mapping':")
	printTypeAndValue(mapping)

}
