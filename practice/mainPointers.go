package main

import (
	"fmt"
)

func mainPointers() {
	var firstName *string = new(string)
	*firstName = "Wilmer"
	fmt.Println(*firstName)
	ptr1 := &firstName
	fmt.Println(ptr1, *ptr1)

	lastName := "Rivera"
	fmt.Println(lastName)
	ptr2 := &lastName
	fmt.Println(ptr2, *ptr2)
	lastName = "Morter"
	fmt.Println(ptr2, *ptr2)

}
