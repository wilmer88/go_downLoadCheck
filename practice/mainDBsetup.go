package main

import (
	"database/sql"
	"fmt"
	"log"
	// "github.com/microsoft/go-mssqldb"
	"net/http"
	"github.com/wilmer88/go_downLoadCheck/controllers"
)
var db11 *sql.DB

func mainSetup() {
	ConnString := "server=localhost;user id=root;passwork=morter706;port=3000;database=wilmer_family"

	var err error
	db11, err = sql.Open("mysql", ConnString)
	if err != nil{
		log.Fatal(err)
	}

	defer db11.Close()

	pingErr := db11.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)}
		fmt.Println("Connected")

	//  creates http server
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
