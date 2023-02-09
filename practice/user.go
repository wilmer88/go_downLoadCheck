package main

import (
	"fmt"
	"github.com/wilmer88/go_downLoadCheck/models"
)

func mainUser() {

	users := models.User{
		ID:        11,
		FirstName: "doris",
		LastName:  "morter",
	}
	fmt.Println(users)

}
