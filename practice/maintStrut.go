package main

import (
	"fmt"
)

func mainStrut() {
	type user struct {
		id        int
		FirstName string
		LastName  string
	}
	var u user
	u.id = 24
	u.FirstName = "wilmer"
	u.LastName = "rivera"
	fmt.Println(u)

	u2 := user{id: 10,
		FirstName: "felix",
		LastName:  "gracia",
	}
	fmt.Println(u2)

}
