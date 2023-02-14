package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wilmer88/go_downLoadCheck/routes"
)

func main() {
	r := mux.NewRouter()
	routes.registerFamilyMemRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000",r))
}