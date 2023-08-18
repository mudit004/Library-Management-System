package models

import (
	"fmt"

	"github.com/tawesoft/golib/v2/dialog"
)

func RemoveBook(title string) error {
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()

	exist, err := BookExists(db, title)
	if err != nil {
		return err
	}

	if exist {
		query := "Select bookID from books where BookName =?"
		var BID int
		err := db.QueryRow(query, title).Scan(&BID)
		if err != nil {
			return err
		}
		query = "Select status from request where BookID=?"
		rows, err := db.Query(query, BID)
		if err != nil {
			return err
		}
		var statusCode []int
		for rows.Next() {
			var status int
			err = rows.Scan(&status)
			if err != nil {
				fmt.Printf("Error %s in scanning row", err)
			}
			statusCode = append(statusCode, status)
		}
		var issueStatus bool = false
		for i := 0; i < cap(statusCode); i++ {
			if statusCode[i] == 1 {
				issueStatus = true
				break
			}
		}

		if issueStatus {
			dialog.Alert("Book is issued. Can't Delete!!!")
			return nil
		}
		query = "Delete from request where BookID=?"
		_, err = db.Exec(query, BID)

		if err != nil {
			return err
		}

		query = "Delete from books where bookID=?"
		_, err = db.Exec(query, BID)

		if err != nil {
			return err
		}
		return nil
	} else {
		dialog.Alert("Book Does Not Exist")
		return nil
	}

}
