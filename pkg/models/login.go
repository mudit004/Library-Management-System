package models

import (
	"database/sql"
	"fmt"
	"lms/pkg/types"
	"lms/pkg/utils"
)

func Login(username, password string) (verify bool, token string, isAdmin bool) {
	db, error := Connection()
	if error != nil {
		fmt.Println(("Database not connected"))
		return
	}
	query := "Select UserID, UName, Hash, Admin from user where Uname=(?) Limit 1"
	var user types.User
	err := db.QueryRow(query, username).Scan(&user.UID, &user.UName, &user.Hash, &user.Admin)
	fmt.Println(user)
	if err != nil {
		fmt.Println("1")
		return false, "", false
	}

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the specified username.")
			return false, "", false
		} else {
			fmt.Println("2")
			return false, "", false

		}
	} else {
		hashPwd := utils.Hash(password)
		fmt.Println(hashPwd)
		check := utils.Check(password, user.Hash)

		if check {
			token := utils.NewJWT(user)
			// Cookie := http.Cookie{
			// 	Name:    "access-token",
			// 	Value:   token,
			// 	Expires: time.Now().Add(48 * time.Hour),
			// 	Path:    "/",
			// }
			var isAdmin bool
			if user.Admin == 1 {
				isAdmin = true
			} else {
				isAdmin = false
			}
			return true, token, isAdmin
		} else {
			fmt.Println("Wrong Password")
			return false, "", false
		}

		// http.SetCookie(w, &cookie)
		// if isAdmin {
		// 	http.Redirect(w, r, "/adminDashboard", http.StatusSeeOther)
		// } else {
		// 	http.Redirect(w, r, "/userDashboard", http.StatusSeeOther)
		// }

	}

}
