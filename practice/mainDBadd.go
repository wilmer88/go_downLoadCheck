package main

import (
	"database/sql"
	"fmt"
	"log"
	// "github.com/microsoft/go-mssqldb"
	"net/http"
	"github.com/wilmer88/go_downLoadCheck/controllers"
)
var db1 *sql.DB

func mainSS() {
	ConnString := "server=localhost;user id=root;passwork=morter706;port=3000;database=wilmer_family"

	var err error
	db, err = sql.Open("mysql", ConnString)
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)}
		fmt.Println("Connected")

	//  creates http server
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
