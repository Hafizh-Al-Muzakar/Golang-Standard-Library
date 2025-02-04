package main

import (
	"container/list"
	"fmt"
)

func main() {

	// creating a new list
	mylist := list.New()

	// adding elements to the list
	mylist.PushBack("Apple")
	mylist.PushBack("Banana")
	mylist.PushBack("Cherry")

	// iterating over the list
	fmt.Println("List Elements:")
	for element := mylist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// removing an element
	elementToRemove := mylist.Front().Next()
	mylist.Remove(elementToRemove)

	// iterating over updated list
	fmt.Println("List Elements:")
	for element := mylist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

}
