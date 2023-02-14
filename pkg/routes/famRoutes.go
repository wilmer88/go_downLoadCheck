package routes

import(
	"github.com/gorilla/mux"
	"github.com/wilmer88/go_downLoadCheck/pkg/controllers"
)

var registerFamilyMemRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.GetUser).Methods("DELETE")
}