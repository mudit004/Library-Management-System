package models

import (
	"fmt"
	"lms/pkg/utils"

	"github.com/tawesoft/golib/v2/dialog"
)

func Register(username, password string) error {
	db, err := Connection()
	if err != nil {
		fmt.Println(("Database not connected"))
		return err
	}

	var userExist int
	err = db.QueryRow("Select Count(*) from user where Uname=?", username).Scan(&userExist)

	if err != nil {
		fmt.Println("Error in retrieving data from User table")
		return err
	}
	if userExist != 0 {
		dialog.Alert("Username already exist")
		return err
	}
	hashedPwd := utils.Hash(password)

	_, err = db.Exec("Insert into user (Uname, hash, admin) Values (?,?,0)", username, hashedPwd)

	if err != nil {
		fmt.Println("Inserting values into user table failed")
		return err
	}
	return nil
}
