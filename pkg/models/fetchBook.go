package models

import (
	"lms/pkg/types"
)

func FetchBook(UserID interface{}) (types.ListBook, error) {

	var nullListBooks types.ListBook
	nullListBooks.Books = nil
	nullListBooks.Requests = nil
	nullListBooks.TotalBooks = 0
	nullListBooks.DifferentBook = 0
	nullListBooks.TotalUsers = 0
	nullListBooks.IssuedBook = 0
	nullListBooks.IssuedAmount = nil
	nullListBooks.UserRequest = nil

	db, err := Connection()
	if err != nil {
		return nullListBooks, err
	}

	selectSQL := "Select * from books"
	rows, err := db.Query(selectSQL)
	defer db.Close()

	if err != nil {
		return nullListBooks, err
	}

	var fetchBook []types.Book

	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.Name, &book.Quantity, &book.Issued)
		if err != nil {
			return nullListBooks, err
		}
		fetchBook = append(fetchBook, book)

	}
	reqSQL := "Select request.requestID, request.UserID, request.BookID, request.status, books.BookName from request Inner join books on request.BookID=books.bookID where (status=0 or status=2)"
	rows, err = db.Query(reqSQL)
	if err != nil {
		return nullListBooks, err
	}
	var fetchRequest []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.Status, &request.BookName)
		if err != nil {
			return nullListBooks, err
		}
		fetchRequest = append(fetchRequest, request)

	}
	reqSQL = "Select request.requestID, request.UserID, request.BookID, request.status, books.BookName from request Inner join books on request.BookID=books.bookID where status=1 AND UserID=?"
	rows, err = db.Query(reqSQL, UserID)
	if err != nil {
		return nullListBooks, err
	}
	var issuedAmount []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.Status, &request.BookName)
		if err != nil {
			return nullListBooks, err
		}
		issuedAmount = append(issuedAmount, request)

	}

	reqSQL = "Select request.requestID, request.UserID, request.BookID, request.status, books.BookName from request Inner join books  on request.BookID=books.bookID where (status=0 or status=2) and request.UserID=?"
	rows, err = db.Query(reqSQL, UserID)
	if err != nil {
		return nullListBooks, err
	}

	var userRequest []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.Status, &request.BookName)
		if err != nil {
			return nullListBooks, err
		}
		userRequest = append(userRequest, request)

	}

	var totalBooks int
	var differentBooks int
	var issuedBooks int
	var totalUsers int

	err = db.QueryRow("Select SUM(quantity), Count(Distinct BookName) from books").Scan(&totalBooks, &differentBooks)

	if err != nil {
		return nullListBooks, err
	}

	err = db.QueryRow("Select Count(*) from request where status =1").Scan(&issuedBooks)
	if err != nil {
		return nullListBooks, err
	}

	err = db.QueryRow("Select Count(*) from user").Scan(&totalUsers)
	if err != nil {
		return nullListBooks, err
	}

	var listBooks types.ListBook
	listBooks.Books = fetchBook
	listBooks.Requests = fetchRequest
	listBooks.TotalBooks = totalBooks
	listBooks.DifferentBook = differentBooks
	listBooks.TotalUsers = totalUsers
	listBooks.IssuedBook = issuedBooks
	listBooks.IssuedAmount = issuedAmount
	listBooks.UserRequest = userRequest
	return listBooks, nil
}
