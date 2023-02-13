package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var dbD *sql.DB

func mainD() {
	//data souce
	dsn := mysql.Config{
		User:   "root",
		Passwd: "morter706",
		Addr:   "localhost",
		Net:    "tcp",
		DBName: "wilmerfamily",
	}

	//databae handle
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")

	rowsUpdated, err := deleteMemberD(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Member found: %v\n", rowsUpdated)

}

func deleteMemberD(famid int32) (int64, error) {
	// use query() instead of Exec() if you dont want result set as the return
	result, err := db.Exec("DELETE FROM fammember WHERE FamID=?",
		 famid)
	if err != nil {
		return 0, fmt.Errorf("delereMember: %v", err)
	}

	id, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("DeletedMember: %v", err)
	}

	return id, nil
}