package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/wilmer88/go_downLoadCheck/models"
)

// handles two types of resource requests. User collection, Mainipulation of user
type userController struct {
	userIdPattern *regexp.Regexp
}

/*
	ServeHTTP(w http.ResponseWriter, r *http.Request) automatically implements go handler interface see go docs for more info

its job is to recive the HTTP request in, based on the information in the request decied which method bellow to pass that request off to, to actually be processed
*/
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAllUsers(w, r)
		case http.MethodPost:
			uc.postUser(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		// strconv converts the matching string into a numerical data type/ a number
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			// checks if conversion returns an error
			w.WriteHeader(http.StatusNotFound)
		}
		// foward the request on to the methods
		switch r.Method {
		case http.MethodGet:
			uc.getUser(id, w)
		case http.MethodPut:
			uc.PutUpdate(id, w, r)
		case http.MethodDelete:
			uc.deleteUser(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

// getAllUsers method is going to handle retriving all of the users from model layer and returning it back out
func (uc *userController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

/*
	getUser method exepts the id of a single resorce, exept respond writer from servHTTP method, call into the model layer and retrive

that user by id. if one is not found it's going to return out an error. line 30 if the id is found encodeResponsAsJSON method is
called; hover over to see details of this method
*/
func (uc *userController) getUser(id int, w http.ResponseWriter) {
	userPerson, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//takes the User object, and it's going turn it into a JSON representation of that object and return it out to the requester
	encodeResponseAsJSON(userPerson, w)
}

// adds a new user to the user collection
func (uc *userController) postUser(w http.ResponseWriter, r *http.Request) {
	userPerson, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse User object or something else went wrong"))
		return
	}
	userPerson, err = models.AddUser(userPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(userPerson, w)
}

// finds the user by id and updates that user
func (uc *userController) PutUpdate(id int, w http.ResponseWriter, r *http.Request) {
	userPerson, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse user object"))
		return
	}
	if id != userPerson.ID {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("id of submitted user must match id in URL"))
		return
	}
	userPerson, err = models.UpdateUser(userPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(userPerson, w)
}

// delets a user with the given id
func (uc userController) deleteUser(idOfUser int, w http.ResponseWriter) {
	err := models.RemoveUserById(idOfUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var userPerson models.User
	err := dec.Decode(&userPerson)
	if err != nil {
		return models.User{}, err
	}
	return userPerson, nil

}

// constructor function with a pointer to the new userConroller object
func newUserController() *userController {
	// provides implementation
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
