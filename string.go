package main

import (
	"fmt"
	"strings"
)

func main() {

	// string declaration
	str := "One Piece"

	// Calculation string lenngth
	length := len(str)
	fmt.Println("String Length:", length)

	// Converting string to uppercase
	strUpper := strings.ToUpper(str)
	fmt.Println("Uppercase:", strUpper)

	// Converting string to lowercase
	strLower := strings.ToLower(str)
	fmt.Println("Lowercase:", strLower)

	// Checking if string contains a substring
	fmt.Println("Contains:", strings.Contains(str, "Piece"))

	// Replacing a substring
	newStr := strings.Replace(str, "Piece", "Punch", -1)
	fmt.Println("Replaced:", newStr)

	// Splitting a string
	Piece := strings.Split(str, " ")
	fmt.Println("Splitted:", Piece)

	// Joining a string
	Joined := strings.Join(Piece, " ")
	fmt.Println("Joined:", Joined)
}
