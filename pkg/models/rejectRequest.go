package models

import (
	"strconv"
)

func RejectRequest(UserID, BookID string) error {
	UserIDint, err := strconv.Atoi(UserID)
	if err != nil {
		return err
	}
	BookIDint, err := strconv.Atoi(BookID)
	if err != nil {
		return err
	}

	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()

	var status int

	_ = db.QueryRow("Select status from request where UID=? and BookID=? ", UserIDint, BookIDint).Scan(&status)

	if status == 0 {
		_, err = db.Exec("Delete from request where UID =? AND BookID=? and status=0", UserIDint, BookIDint)

		if err != nil {
			return err
		}

	} else if status == 2 {
		_, err = db.Exec("Update request set status=1 where UID=? AND BookID=? and status=2", UserIDint, BookIDint)

		if err != nil {
			return err
		}

	}
	return nil
}
