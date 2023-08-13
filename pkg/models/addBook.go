package models

import (
	"database/sql"
	"fmt"
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

func AddBook(title string, quantity string) string {
	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return "Database Connection Error"
	}
	defer db.Close()
	fmt.Println(title, " : ", quantity)
	quantityint, err := strconv.Atoi(quantity)
	if err != nil {
		fmt.Println("Error:", err)
		return "error in parsing"
	}

	exist, err := BookExists(db, title)
	if err != nil {
		return "Command error"
	}
	if quantityint > 0 {
		if exist {
			query := "UPDATE books SET quantity = ? WHERE BookName=?"
			_, err := db.Exec(query, quantityint, title)
			if err != nil {
				panic(err)
			}
		} else {
			query := "INSERT INTO books (BookName, quantity) VALUES (?, ?)"
			_, err = db.Exec(query, title, quantityint)
			if err != nil {
				fmt.Println(err)
				return "Unexpected Error occured"
			}
		}

		fmt.Println("Operation successful")
		return ""
	} else {
		fmt.Println("Invalid quantity")
		return "Invalid quantity"
	}

}
