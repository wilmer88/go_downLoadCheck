package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/wilmer88/go_downLoadCheck/controllers"
)
var db *sql.DB

func main() {
	dsn := mysql.Config{
		User: "root",
		Passwd: "morter706",
		Net: "tcp",
		Addr: "localhost:3306",
		DBName: "wilmer_family",

	}

	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
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
