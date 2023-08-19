package models

import (
	"database/sql"
	"fmt"

	"github.com/tawesoft/golib/v2/dialog"
)

func RequestExists(db *sql.DB, UserID interface{}, BookID int) (bool, error) {
	query := "SELECT COUNT(*) FROM request WHERE UserID=? and BookID=? and status<>3"
	var count int
	err := db.QueryRow(query, UserID, BookID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func RequestBook(UserID interface{}, BookID int) error {
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()
	var quantity int
	err = db.QueryRow("Select quantity from books where BookID=?", BookID).Scan(&quantity)
	if err != nil {
		return err
	}
	if quantity > 0 {
		req, err := RequestExists(db, UserID, BookID)
		if err != nil {
			return err
		}
		if !req {
			_, err = db.Exec("Insert Into request (UserID, BookID, status) Values (?, ?, 0)", UserID, BookID)
			if err != nil {
				return err
			}

			return nil
		} else {
			dialog.Alert("Already requested for book")
			return nil
		}
	} else {
		dialog.Alert("Book is Not Available!!")
		return fmt.Errorf("book is not available")
	}

}
