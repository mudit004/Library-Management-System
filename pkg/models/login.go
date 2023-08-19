package models

import (
	"database/sql"
	"lms/pkg/types"
	"lms/pkg/utils"
)

func Login(username, password string) (verify bool, token string, isAdmin bool, errors error) {
	db, err := Connection()
	if err != nil {
		return false, "", false, err
	}
	query := "Select UserID, UserName, Hash, Admin from user where UserName=(?) Limit 1"
	var user types.User
	err = db.QueryRow(query, username).Scan(&user.UserID, &user.UserName, &user.Hash, &user.Admin)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", false, nil
		} else {
			return false, "", false, err

		}
	} else {
		check := utils.Check(password, user.Hash)

		if check {
			token, err := utils.NewJWT(user)
			if err != nil {
				return false, "", false, err
			}
			var isAdmin bool
			if user.Admin == 1 {
				isAdmin = true
			} else {
				isAdmin = false
			}
			return true, token, isAdmin, nil
		} else {
			return false, "", false, nil
		}

	}

}
