package service_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kakts/go-orm-sbx/gorm_app/service"
)

// TestGetUserByName a successful case for GetUserByName
func TestGetUserByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userService := service.Service{}

	userModel := userService.GetUserByName("notFound")
	print(len(userModel))
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
