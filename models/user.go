package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(userPerson User) (User, error) {
	userPerson.ID = nextId
	nextId++
	users = append(users, &userPerson)
	return userPerson, nil
}
