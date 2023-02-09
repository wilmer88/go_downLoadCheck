package controllers

import (
	"net/http"
	"regexp"
)
// handles two types of resource requests. User collection, Mainipulation of user
type userController struct {
	userIdPattern *regexp.Regexp
}

// ServeHTTP(w http.ResponseWriter, r *http.Request) automatically implements go handler interface see go docs for more info
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("I come from the userController"))

}

// constructor function with a pointer to the new userConroller object
func newUserController() *userController {
	// provides implementation
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}