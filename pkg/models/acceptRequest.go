package models

import (
	"fmt"
	"strconv"
)

func AcceptRequest(UserID, BookID string) string {
	UserIDint, err := strconv.Atoi(UserID)
	if err != nil {
		fmt.Println("Error:", err)
		return "error in parsing"
	}
	BookIDint, err := strconv.Atoi(BookID)
	if err != nil {
		fmt.Println("Error:", err)
		return "error in parsing"
	}
	fmt.Println(UserIDint, " : ", BookIDint)

	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return "Database Connection Error"
	}
	defer db.Close()

	var status int

	_ = db.QueryRow("Select status from request where UID=? and BookID=? and status<>1 and status<>3", UserIDint, BookIDint).Scan(&status)

	if status == 0 {
		fmt.Println("Updating 0 to 1")
		_, err = db.Exec("Update request set status = 1 where UID=? and BookID=? and status = 0", UserIDint, BookIDint)

		if err != nil {
			fmt.Println("Error in Updating table")
			return "error in updating table"
		}
	} else if status == 2 {
		fmt.Println("Updating 2 to 3")

		_, err = db.Exec("Update request set status = 3 where UID=? and BookID=? and status =2", UserIDint, BookIDint)

		if err != nil {
			fmt.Println("Error in Updating table")
			return "Error in updating table"
		}

		_, err = db.Exec("Update books set quantity=quantity-1 where bookID=?", BookIDint)

		if err != nil {
			fmt.Println("Error in Updating table")
			return "Error in updating table"
		}
	}
	return ""
}
