package controllers

import (
	"net/http"
	"regexp"
	"github.com/wilmer88/go_downLoadCheck/models"
)

// handles two types of resource requests. User collection, Mainipulation of user
type userController struct {
	userIdPattern *regexp.Regexp
}

// ServeHTTP(w http.ResponseWriter, r *http.Request) automatically implements go handler interface see go docs for more info
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, I come from the userController"))
}
// getAll method is going to handle retriving all of the users from model layer and returning it back out
func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
 		encodeResponseAsJSON(models.GetUsers(), w)
}
// 
func(uc *userController) get(id int, w http.ResponseWriter) {
	userPerson, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(userPerson, w)
}


// constructor function with a pointer to the new userConroller object
func newUserController() *userController {
	// provides implementation
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}