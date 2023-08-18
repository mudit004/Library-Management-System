package models

import (
	"database/sql"
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

func ReturnBook(UID interface{}, BID int) error {
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()

	status, err := BookIssued(db, UID, BID)
	if err != nil {
		return err
	}
	if status {
		_, err = db.Exec("Update request set status = 2 where UID =? and BookID = ? and status=1", UID, BID)
		if err != nil {
			return err
		}

	}
	return nil

}
