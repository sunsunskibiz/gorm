package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sun?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Drop table
	db.Migrator().DropTable(&User{}, &CreditCard{}, &Image{})

	// Migrate the schema
	db.AutoMigrate(&User{}, &CreditCard{}, &Image{})

	// Create
	db.Create(&User{Name: "MU", UserID: "1"})
	db.Create(&CreditCard{Number: "111", UserID: "1", CreditCardID: "1"})
	db.Create(&CreditCard{Number: "222", UserID: "1", CreditCardID: "2"})
	db.Create(&Image{Url: "url-one-one", CreditCardID: "1"})
	db.Create(&Image{Url: "url-one-two", CreditCardID: "1"})

	// Read
	var user User
	db.First(&user, 1)
	fmt.Printf("User: %+v\n", user)

	users, err := GetAll(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Users: %+v\n", users)

	result, err := GetResult(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %+v\n", result)
}
