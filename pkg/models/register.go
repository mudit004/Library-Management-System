package models

import (
	"lms/pkg/utils"

	"github.com/tawesoft/golib/v2/dialog"
)

func Register(username, password string) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	var userExist int
	err = db.QueryRow("Select Count(*) from user where Username=?", username).Scan(&userExist)

	if err != nil {
		return err
	}
	if userExist != 0 {
		dialog.Alert("Username already exist")
		return nil
	}
	hashedPassword, err := utils.Hash(password)

	if err != nil {
		return err
	}

	_, err = db.Exec("Insert into user (Username, hash, admin) Values (?,?,0)", username, hashedPassword)

	if err != nil {
		return err
	}
	return nil
}
