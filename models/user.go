package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// variable block. creates a collection of users in a/as a slice, that hold pointers to user objects
var (
	users []*User
	// holds an id indexer for this app, normally a primary key that comes from a data base created
	nextId = 1
	// nextId int32 = 1
)

// returns all User collection
func GetUsers() []*User {
	return users
}

// adds User to the User collection. requires a user object  and returns the user that was created and a potential error object
func AddUser(userPerson User) (User, error) {
	userPerson.ID = nextId
	nextId++
	users = append(users, &userPerson)
	return userPerson, nil
}
