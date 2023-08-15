package models

import (
	"database/sql"
	"fmt"
)

func BookIssued(db *sql.DB, UID interface{}, BID int) (bool, error) {
	query := "SELECT COUNT(*) FROM request WHERE UID=? and BookID=? and status=1"
	var count int
	err := db.QueryRow(query, UID, BID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func ReturnBook(UID interface{}, BID int) string {
	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return "Database Connection Error"
	}
	defer db.Close()

	status, _ := BookIssued(db, UID, BID)
	fmt.Println(status)
	if status {
		_, err = db.Exec("Update request set status = 2 where UID =? and BookID = ? and status=1", UID, BID)
		if err != nil {
			panic(err)
		}

	}
	return ""

}
