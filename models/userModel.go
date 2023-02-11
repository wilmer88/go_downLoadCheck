package models

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID        int
	FirstName string
	Happiness int
	Url 		string
	CreatedAt time.Time
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
	if userPerson.ID != 0 {
		return User{}, errors.New("id was provided not expecting id")
	}
	userPerson.ID = nextId
	nextId++
	userPerson.CreatedAt = time.Now()
	users = append(users, &userPerson)
	return userPerson, nil
}

func GetUserByID(PersonId int) (User, error) {
	for _, userPerson := range users {
		if userPerson.ID == PersonId {
			return *userPerson, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", PersonId)

}

func UpdateUser(userPerson User) (User, error) {
	for i, toUpdate1 := range users {
		if toUpdate1.ID == userPerson.ID {
			users[i] = &userPerson
			return userPerson, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' was not found", userPerson.ID)
}

func RemoveUserById(userPersonId int) error {
	for i, userPerson := range users {
		if userPerson.ID == userPersonId {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' to delete was not found", userPersonId)
}
