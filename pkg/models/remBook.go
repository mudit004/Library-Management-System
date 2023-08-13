package models

import "fmt"

func RemBook(title string) string {
	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return "Database Connection Error"
	}
	defer db.Close()

	exist, err := BookExists(db, title)
	if err != nil {
		return "BookExists function Error"
	}

	if exist {
		query := "Select bookID from books where BookName =?"
		var BID int
		err := db.QueryRow(query, title).Scan(&BID)
		if err != nil {
			fmt.Println("Error in extracting data")
		}
		query = "Delete from request where BookID=?"
		_, err = db.Exec(query, BID)

		if err != nil {
			return err.Error()
		}

		query = "Delete from books where bookID=?"
		_, err = db.Exec(query, BID)

		if err != nil {
			return err.Error()
		}
		return ""
	} else {
		return "Book Does Not Exist"
	}

}
