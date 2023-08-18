package models

import (
	"strconv"
)

func AcceptRequest(UserID, BookID string) error {
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

	_ = db.QueryRow("Select status from request where UID=? and BookID=? and status<>1 and status<>3", UserIDint, BookIDint).Scan(&status)

	if status == 0 {
		_, err = db.Exec("Update request set status = 1 where UID=? and BookID=? and status = 0", UserIDint, BookIDint)

		if err != nil {
			return err
		}
		_, err = db.Exec("Update books set quantity=quantity-1 where bookID=?", BookIDint)

		if err != nil {
			return err
		}

		_, err = db.Exec("Update books set issued=issued+1 where bookID=?", BookIDint)

		if err != nil {
			return err
		}
	} else if status == 2 {

		_, err = db.Exec("Delete from request where status = 2 and UID=? and BookID=?", UserIDint, BookIDint)

		if err != nil {
			return err
		}

		_, err = db.Exec("Update books set quantity=quantity+1 where bookID=?", BookIDint)

		if err != nil {
			return err
		}
		_, err = db.Exec("Update books set issued=issued-1 where bookID=?", BookIDint)

		if err != nil {
			return err
		}

	}
	return nil
}
