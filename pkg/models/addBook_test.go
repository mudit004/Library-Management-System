package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestBookExists(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	count := 1
	bookName := "tester"
	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	mock.ExpectQuery("SELECT COUNT(.*) FROM books WHERE BookName = ?").WithArgs(bookName).WillReturnRows(rows)

	result, err := BookExists(db, bookName)

	if err != nil {
		t.Errorf("Error faced: %s and result is: %t", err, result)
		t.Fatal("Error in Function")
	}

	expected := true

	if result != expected {
		t.Errorf("got %v, wanted %v", result, expected)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
