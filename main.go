package main

import (
	
	"github.com/wilmer88/go_downLoadCheck/models"
	"fmt"

)



func main() {

	u := models.User{
		ID: 2,
		FirstName: "doris",
		LastName: "morter",

	}
	fmt.Println(u)


}