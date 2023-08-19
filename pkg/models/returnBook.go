package models

import (
	"database/sql"
)

func BookIssued(db *sql.DB, UserID interface{}, BookID int) (bool, error) {
	query := "SELECT COUNT(*) FROM request WHERE UserID=? and BookID=? and status=1"
	var count int
	err := db.QueryRow(query, UserID, BookID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func ReturnBook(UserID interface{}, BookID int) error {
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()

	status, err := BookIssued(db, UserID, BookID)
	if err != nil {
		return err
	}
	if status {
		_, err = db.Exec("Update request set status = 2 where UserID =? and BookID = ? and status=1", UserID, BookID)
		if err != nil {
			return err
		}

	}
	return nil

}
