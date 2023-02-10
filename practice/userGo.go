package main
import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string

}

func (u Person) String() string{
	return fmt.Sprintf(
		"First Name:\t\t%q\n"+
		"Last Name: \t\t%q\n"+
		"Time Created: \t%v", u.FirstName, u.LastName)
}

var persons = []Person {
	Person{
		ID: 1,
		FirstName: "wilmer",
		LastName: "rivera",
	},
	Person{
		ID: 2,
		FirstName: "felix",
		LastName: "gracia",
	},
	Person{
		ID: 3,
		FirstName: "doris",
		LastName: "morter",
	}}