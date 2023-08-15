package models

import (
	"fmt"

	"github.com/tawesoft/golib/v2/dialog"
)

func RemoveBook(title string) string {
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
			return "Error in extracting data"
		}
		query = "Select status from request where BookID=?"
		rows, err := db.Query(query, BID)
		var statusCode []int
		for rows.Next() {
			var status int
			err = rows.Scan(&status)
			if err != nil {
				fmt.Printf("Error %s in scanning row", err)
			}
			statusCode = append(statusCode, status)
		}
		fmt.Println("------------>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>----------------------------------------------->>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ", statusCode)
		var issueStatus bool = false
		for i := 0; i < cap(statusCode); i++ {
			if statusCode[i] == 1 {
				issueStatus = true
				break
			}
		}

		if issueStatus {
			fmt.Println("Hello Guyzz <<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ")
			dialog.Alert("Book is issued. Can't Delete!!!")
			return "Books Already Issued"
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
