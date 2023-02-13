package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var dbU *sql.DB

func mainU() {
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

	rowsUpdated, err := updateMemberU("fany", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Member found: %v\n", rowsUpdated)

}

func updateMemberU(firstname string, famid int32) (int64, error) {
	// use query() instead of Exec() if you dont want result set as the return
	result, err := db.Exec("UPDATE fammember SET FirstName=? WHERE FamID=?",
		firstname, famid)
	if err != nil {
		return 0, fmt.Errorf("updateMember: %v", err)
	}

	id, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("updateMember: %v", err)
	}

	return id, nil
}
