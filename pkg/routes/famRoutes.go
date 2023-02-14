package routes

import(
	"github.com/gorilla/mux"
	"github.com/wilmer88/go_downLoadCheck/pkg/controllers"
)

var RegisterFamilyMemRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreatefamiliaMiembro).Methods("POST")
	router.HandleFunc("/user", controllers.GetMembers).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetFamMemById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateFamMember).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.EleminaMember).Methods("DELETE")
}