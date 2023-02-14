package main

import (
	"github.com/wilmer88/go_downLoadCheck/pkg/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_"github.com/jinzhu/gorm/dialects/mysql"
	
)

func main() {
	r := mux.NewRouter()
	routes.RegisterFamilyMemRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000",r))
}