package models

import (
	"database/sql"
	"strconv"
)

func BookExists(db *sql.DB, title string) (bool, error) {
	query := "SELECT COUNT(*) FROM books WHERE BookName = ?"
	var count int
	err := db.QueryRow(query, title).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddBook(title string, quantity string) error {
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()
	quantityint, err := strconv.Atoi(quantity)
	if err != nil {
		return err
	}

	exist, err := BookExists(db, title)
	if err != nil {
		return err
	}
	if quantityint > 0 {
		if exist {
			query := "UPDATE books SET quantity = ? WHERE BookName=?"
			_, err := db.Exec(query, quantityint, title)
			if err != nil {
				return err
			}
		} else {
			query := "INSERT INTO books (BookName, quantity) VALUES (?, ?)"
			_, err = db.Exec(query, title, quantityint)
			if err != nil {
				return err
			}
		}

		return nil
	} else {

		return err
	}

}
