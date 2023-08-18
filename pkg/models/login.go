package models

import (
	"database/sql"
	"fmt"
	"lms/pkg/types"
	"lms/pkg/utils"
)

func Login(username, password string) (verify bool, token string, isAdmin bool, errors error) {
	db, err := Connection()
	if err != nil {
		fmt.Println(("Database not connected"))
		return false, "", false, err
	}
	query := "Select UserID, UName, Hash, Admin from user where Uname=(?) Limit 1"
	var user types.User
	err = db.QueryRow(query, username).Scan(&user.UID, &user.UName, &user.Hash, &user.Admin)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", false, nil
		} else {
			return false, "", false, err

		}
	} else {
		// hashPwd := utils.Hash(password)
		// fmt.Println(hashPwd)
		check := utils.Check(password, user.Hash)

		if check {
			token := utils.NewJWT(user)
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
