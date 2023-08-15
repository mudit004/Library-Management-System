package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tawesoft/golib/v2/dialog"
)


func IncrementBook(BookID, amount string) error{
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Error:", err)
		dialog.Alert("Invalid Quantity")
		return err
	}
	BookIDint,_ := strconv.Atoi(BookID)

	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return err
	}
	defer db.Close()

	query:= "update books set quantity = quantity+ ? where BookID=?"

	_,err = db.Query(query,amountInt,BookIDint)

	if err != nil {
		return err
	}


	return nil
}

func DecrementBook(BookID, amount string) error{
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	BookIDint, err := strconv.Atoi(BookID)
	db, err := Connection()

	if err != nil {
		fmt.Println("Error in Connecting to database")
		return err
	}
	defer db.Close()

	var quantity int
	err = db.QueryRow("Select quantity from books where BookID=?", BookIDint).Scan(&quantity)
	if err != nil {
		return err
	}

	if quantity>=amountInt{

	query:= "update books set quantity = quantity- ? where BookID=?"

	_,err = db.Query(query,amountInt,BookIDint)

	if err != nil {
		return err
	}


	return nil
	} else{
		dialog.Alert("number of books are less than you want to decrease!!")
		err := errors.New("number of books less")
		return err
	}
}