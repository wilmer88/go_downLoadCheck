package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/wilmer88/go_downLoadCheck/controllers"
)
var db1 *sql.DB

func mainCrud() {
	dsn := mysql.Config{
		User: "root",
		Passwd: "morter706",
		Addr: "localhost:3306",
		DBName: "wilmer_family",}

	var err error
	db1, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil{
		log.Fatal(err)}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)}
		fmt.Println("Connected")
		
		memberID, err := addMember1(1,"doris", 5)
		if err != nil {
			log.Fatal(err)
		}
	
	fmt.Printf("ID of added member: %v\n", memberID)
	//  creates http server
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}

func addMember1(ID int, FirstName string, Happiness int) (int64, error) {
	result, err := db.Exec("INSERT INTO FamMember (ID, FirstName, Happiness) VALUES(?,?,?)",
ID,FirstName, Happiness)
if err != nil {
	return 0, fmt.Errorf("addMember: %v", err)
}
id, err := result.LastInsertId()
if err != nil {
	return 0, fmt.Errorf("addMember: %v", err)
}
return id,nil
}
// C:\code\firstGo