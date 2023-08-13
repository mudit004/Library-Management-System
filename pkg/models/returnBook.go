package models

import (
	"database/sql"
	"fmt"
)

func BookIssued(db *sql.DB, UID, BID int) (bool, error) {
	query := "SELECT COUNT(*) FROM request WHERE UID=? and BookID=? and status=1"
	var count int
	err := db.QueryRow(query, UID, BID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func ReturnBook(UID, BID int) string {
	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return "Database Connection Error"
	}
	defer db.Close()

	status, _ := BookIssued(db, UID, BID)

	if status {
		_, err = db.Exec("Update requeat set status = 3 where UID =? and BookID = ?", UID, BID)
		if err != nil {
			panic(err)
		}

	}
	return ""

}
