package models

import (
	"fmt"
	"lms/pkg/types"
)

func FetchBook(UID interface{}) types.ListBook {
	UIDint, _ := UID.(int)
	fmt.Println(UID, "---------------", UIDint)
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to database")
	}

	selectSQL := "Select * from books"
	rows, err := db.Query(selectSQL)
	fmt.Println("fetchbook", rows)
	defer db.Close()

	if err != nil {
		fmt.Println("Error in Retrieving data for html books template")
	}

	var fetchBook []types.Book

	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.Name, &book.Quantity)
		if err != nil {
			fmt.Printf("Error %s in scanning row", err)
		}
		fetchBook = append(fetchBook, book)

	}
	fmt.Println("fetchBOOK RESULT", fetchBook)
	reqSQL := "Select request.reqID, request.UID, request.BookID, request.status, books.BookName from request Inner join books on request.BookID=books.bookID where (status=0 or status=2)"
	rows, err = db.Query(reqSQL)
	if err != nil {
		fmt.Println("Error in Retrieving data for html books template")
	}
	fmt.Println("for fetchRequest", rows)
	var fetchRequest []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.Status, &request.BookName)
		if err != nil {
			fmt.Printf("Error %s in scanning row", err)
		}
		fetchRequest = append(fetchRequest, request)

	}
	fmt.Println("fetchRequest RESULT", fetchRequest)
	reqSQL = "Select request.reqID, request.UID, request.BookID, request.status, books.BookName from request Inner join books on request.BookID=books.bookID where status=1 AND UID=?"
	rows, err = db.Query(reqSQL, UID)
	fmt.Println("for issued", rows)
	if err != nil {
		fmt.Println("Error in Retrieving data for html books template")
	}
	var issued []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.Status, &request.BookName)
		if err != nil {
			fmt.Printf("Error %s in scanning row", err)
		}
		fmt.Println(request)
		issued = append(issued, request)

	}
	fmt.Println("issued Result", issued)
	fmt.Println("UID :", UIDint)
	reqSQL = "Select request.reqID, request.UID, request.BookID, request.status, books.BookName from request Inner join books  on request.BookID=books.bookID where (status=0 or status=2) and request.UID=?"
	rows, err = db.Query(reqSQL, UID)
	if err != nil {
		fmt.Println("Error in Retrieving data for html books template")
	}

	var userRequest []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.Status, &request.BookName)
		if err != nil {
			fmt.Printf("Error %s in scanning row", err)
		}
		userRequest = append(userRequest, request)

	}
	var totalBooks int
	var diffBooks int
	var issuedBooks int
	var totalUsers int

	err = db.QueryRow("Select SUM(quantity), Count(Distinct BookName) from books").Scan(&totalBooks, &diffBooks)

	err = db.QueryRow("Select Count(*) from request where status =1").Scan(&issuedBooks)

	err = db.QueryRow("Select Count(*) from user").Scan(&totalUsers)

	var listBooks types.ListBook
	listBooks.Books = fetchBook
	listBooks.Requests = fetchRequest
	listBooks.TotalBooks = totalBooks
	listBooks.DiffBook = diffBooks
	listBooks.TotalUsers = totalUsers
	listBooks.IssuedBook = issuedBooks
	listBooks.Issued = issued
	listBooks.UserRequest = userRequest
	fmt.Println(listBooks)
	return listBooks
}
