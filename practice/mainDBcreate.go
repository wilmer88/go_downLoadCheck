package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql"


)
var dbCreate *sql.DB

func mainCreate() {
	dsn := mysql.Config{
		User: "root",
		Passwd: "morter706",
		Addr: "localhost",
		DBName: "wilmerfamily",}

	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil{
		log.Fatal(err)}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)}
		fmt.Println("Connected")
		
		memberID, err := addMember1(3,"doris", 5)
		if err != nil {
			log.Fatal(err)
		}
	
	fmt.Printf("ID of added member: %v\n", memberID)
	//  creates http server


}

func addMember1(FamID int, FirstName string, Happiness int) (int64, error) {
	result, err := db.Exec("INSERT INTO fammember (FamID, FirstName, Happiness) VALUES(?,?,?)",
FamID,FirstName, Happiness)
if err != nil {
	return 0, fmt.Errorf("addMember: %v", err)
}
id, err := result.LastInsertId()
if err != nil {
	return 0, fmt.Errorf("addMember: %v", err)
}
return id,nil
}