package models

import (
	"fmt"
	"lms/pkg/types"
)

func FetchBook() types.ListBook {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to database")
	}

	selectSQL := "Select * from books"
	rows, err := db.Query(selectSQL)

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

	reqSQL := "Select * from request where status =0 or status= 2"
	rows, err = db.Query(reqSQL)
	if err != nil {
		fmt.Println("Error in Retrieving data for html books template")
	}

	var fetchRequest []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.Status)
		if err != nil {
			fmt.Printf("Error %s in scanning row", err)
		}
		fetchRequest = append(fetchRequest, request)

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
	fmt.Println(listBooks)
	return listBooks
}
