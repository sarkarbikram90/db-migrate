package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDatabaseMigration(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create a new GORM database instance using the mock database connection
	_, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	// Set up expectations for the auto-migrate call
	mock.ExpectExec("CREATE TABLE `User`").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("CREATE TABLE `Product`").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("CREATE TABLE `ProductPurchased`").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("CREATE TABLE `ContactInfo`").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("CREATE TABLE `PurchaseTime`").WillReturnResult(sqlmock.NewResult(0, 1))

	// Call the main function which triggers database migration
	// Test the main function
	t.Run("Main", func(t *testing.T) {
		// Call the main function which triggers database migration
		main()
	})
	if err != nil {
		t.Fatalf("Error migrating schemas: %v", err)
	}

	// Assert that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
