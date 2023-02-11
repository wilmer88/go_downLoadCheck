/*
	this file handles rought routing

when a  network request recived

	will direct to the correct controller to be processed
*/
package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterController responsibility is to create a newUserController
func RegisterControllers() {
	// newuserController is a constructor function
	uc := newUserController()
	// handeles http.handle with users without extra information
	http.Handle("/users", *uc)
	// handles http.handle with users with extra info
	http.Handle("/users/", *uc)
}

/*
	Reaches into encoding/json package, creates an encoder that is designed to encode Go objects into JSON representations,

then calls that Encod method on that encoder/ enc passing on whatever data recived
*/
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
