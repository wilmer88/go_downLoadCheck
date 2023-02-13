package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func mainRead() {
	//data souce
	dsn := mysql.Config{
		User: "root",
		Passwd: "morter706",
		Addr: "localhost:3306",
		Net: "tcp",
		DBName: "wilmer_family",
	}

//databae handle
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)
	}
		fmt.Println("Connected")

		FamMember, err := GetMember1(3)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Member found: %v\n", FamMember)
	}

type Member struct {
	ID int
	FirstName string
	Happiness int
}

func GetMember1(memberID int32) ([]Member,error) {
		var members []Member

		result, err := db.Query("SELECT ID,	FirstName, Happiness FROM FamMember WHERE ID= ?", 
		memberID)
		if err != nil{
			return nil, fmt.Errorf("getMember %v: %v", memberID, err )}

		defer result.Close()
	
		for result.Next() {
			var isFam Member
			if err := result.Scan(&isFam.ID, &isFam.FirstName, &isFam.Happiness); err != nil {
				return nil, fmt.Errorf("getMember %v: %v", memberID, err)}

			members = append(members, isFam)
			if err:= result.Err(); err != nil {
				return nil , fmt.Errorf("getMember %v: %v", memberID, err)}}
		return members,nil
	}

