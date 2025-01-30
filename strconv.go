package main

import (
	"fmt"
	"strconv"
)

func main() {

	// convert string to integer
	str := "9"
	numVal, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error Converting String to Integer:", err)
	} else {
		fmt.Println("String to Integer:", numVal)
	}

	// convert integer to string
	numVal = 9
	str = strconv.Itoa(numVal)
	fmt.Println("Integer to String:", str)

	// convert string to boolean
	strBool := "true"
	boolVal, err := strconv.ParseBool(strBool)
	if err != nil {
		fmt.Println("Error Converting String to Boolean:", err)
	} else {
		fmt.Println("String to Boolean:", boolVal)
	}

	// convert string to float
	strFloat := "3.14"
	floatVal, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Error Converting String to Float:", err)
	} else {
		fmt.Println("String to Float:", floatVal)
	}

	// convert float to string
	floatVal = 3.14
	strFloat = strconv.FormatFloat(floatVal, 'f', -1, 64)
	fmt.Println("Float to String:", strFloat)

}
