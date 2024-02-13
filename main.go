package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Product struct {
	gorm.Model
	Name  string
	Price float64
}

type ProductPurchased struct {
	gorm.Model
	UserID      uint
	ProductID   uint
	Quantity    int
	ProductTime time.Time
}

type ContactInfo struct {
	gorm.Model
	UserID  uint
	Phone   string
	Address string
	City    string
	PINCode string
}

type PurchaseTime struct {
	gorm.Model
	UserID       uint
	ProductID    uint
	PurchaseTime time.Time
}

// Custom table names using tags
func (User) TableName() string {
	return "User"
}

func (Product) TableName() string {
	return "Product"
}

func (ProductPurchased) TableName() string {
	return "ProductPurchased"
}

func (ContactInfo) TableName() string {
	return "ContactInfo"
}

func (PurchaseTime) TableName() string {
	return "PurchaseTime"
}

func main() {
	// Set up database connection
	dsn := "dbusername:dbpassword@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.AutoMigrate(&User{}, &Product{}, &ProductPurchased{}, &ContactInfo{}, &PurchaseTime{})
	if err != nil {
		log.Fatalf("Error migrating schemas: %v", err)
	}

	fmt.Println("Database migration is successful!")
}
