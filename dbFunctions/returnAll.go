package dbFunctions

import (
	"fmt"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
func sendUsers() ([]Member, error) {
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
		return nil, err
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	FamMember, err := GetMember1(1)
	if err != nil {
		return nil, err
	}

	return FamMember, nil
}
type Member struct {
	FamID     int
	FirstName string
	Happiness int
}

func GetMember1(famID int32) ([]Member, error) {
	var members []Member

	result, err := db.Query("SELECT * FROM fammember ORDER BY FamID DESC", famID)
	if err != nil {
		return nil, fmt.Errorf("getMember %v: %v", famID, err)
	}

	defer result.Close()

	for result.Next() {
		var isFam Member
		if err := result.Scan(&isFam.FamID, &isFam.FirstName, &isFam.Happiness); err != nil {
			return nil, fmt.Errorf("getMember %v: %v", famID, err)
		}

		members = append(members, isFam)
		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("getMember %v: %v", famID, err)
		}
	}

	return members, nil
}