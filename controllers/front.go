package controllers

/* this file handles rought routing
when a  network request recived
 will direct to the correct controller to be processed */

import "net/http"

// RegisterController responsibility is to create a newUserController
func RegisterControllers() {

	// newuserController is a constructor function
	uc := newUserController()

	// handeles http.handle with users without extra information
	http.Handle("/users", *uc)
	// handles http.handle with users with extra info
	http.Handle("/users/", *uc)
}
