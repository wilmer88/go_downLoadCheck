package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var dbc *sql.DB

func mainConnect() {
	//data souce
	dsn := mysql.Config{
		User: "root",
		Passwd: "morter706",
		Addr: "localhost:3306",
		Net: "tcp",
		DBName: "wilmerFamily",
	}

//databae handle
	var err error
	dbc, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil{
		log.Fatal(err)
	}

	defer dbc.Close()

	pingErr := dbc.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)
	}
		fmt.Println("Connected")

	}
