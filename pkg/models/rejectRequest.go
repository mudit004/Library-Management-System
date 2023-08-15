package models

import (
	"fmt"
	"strconv"
)

func RejectRequest(UserID, BookID string) string {
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

	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return "Database Connection Error"
	}
	defer db.Close()

	var status int

	err = db.QueryRow("Select status from request where UID=? and BookID=?", UserIDint, BookIDint).Scan(&status)

	if status == 0 {
		_, err = db.Exec("Delete from request where UID =? AND BookID=? and status=0", UserIDint, BookIDint)

		if err != nil {
			fmt.Println("Error in Updating table")
			return "error in updating table"
		}
		_, err = db.Exec("Update books set quantity=quantity+1 where bookID=?", BookIDint)

		if err != nil {
			fmt.Println("Error in Updating table")
			return "Error in updating table"
		}
	} else if status == 2 {
		_, err = db.Exec("Update request set status=1 where UID=? AND BookID=? and status=2", UserIDint, BookIDint)

		if err != nil {
			fmt.Println("Error in Updating table")
			return "Error in updating table"
		}

	}
	return ""
}
