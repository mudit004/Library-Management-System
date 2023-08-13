package models

import (
	"database/sql"
	"fmt"
)

func RequestExists(db *sql.DB, UID, BID int) (bool, error) {
	query := "SELECT COUNT(*) FROM request WHERE UID=? and BookID=? and status<>3"
	var count int
	err := db.QueryRow(query, UID, BID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func RequestBook(UID, BID int) string {
	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return "Database Connection Error"
	}
	defer db.Close()
	var quantity int
	err = db.QueryRow("Select quantity from books where BookID=?", BID).Scan(&quantity)
	if err != nil {
		return "Book Does Not exist"
	}
	if quantity > 0 {
		req, _ := RequestExists(db, UID, BID)
		if !req {
			_, err = db.Exec("Insert Into request (UID, BookID, status) Values (?, ?, 1)", UID, BID)
			if err != nil {
				return "Error in database Insertion"
			}
			_, err = db.Exec("Update books set quantity=quantity-1 where bookID=?", BID)

			if err != nil {
				fmt.Println("Error in Updating table")
				return "Error in updating table"
			}
			return ""
		} else {
			return "Already requested for book"
		}
	} else {
		return "Invalid Quantity"
	}

}
