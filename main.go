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
	db.Create(&User{Name: "MU", ID: "1"})
	db.Create(&CreditCard{Number: "111", UserID: "1", ID: "1"})
	db.Create(&CreditCard{Number: "222", UserID: "1", ID: "2"})
	db.Create(&Image{Url: "url-one-one", CreditCardID: "1", ID: "1"})
	db.Create(&Image{Url: "url-one-two", CreditCardID: "1", ID: "2"})

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

	// Update
	var userWithAsso User
	db.Preload("CreditCards.Images").First(&userWithAsso, 1)
	fmt.Printf("userWithAsso before update: %+v\n", userWithAsso)

	userWithAsso.Name = "pikachu"
	userWithAsso.CreditCards = append(userWithAsso.CreditCards, CreditCard{Number: "333", UserID: "1", ID: "3"})
	db.Updates(&userWithAsso)

	db.Preload("CreditCards.Images").First(&userWithAsso, 1)
	fmt.Printf("userWithAsso after update: %+v\n", userWithAsso)

	// Update user with partial field (just name)
	updateUser := User{
		ID: "1",
		Name: "charmander",
		CreditCards: []CreditCard{{Number: "444", UserID: "1", ID: "4"}},
	}
	fmt.Printf("updateUser before update: %+v\n", updateUser)
	
	db.Updates(&updateUser)

	// TODO: Remove assoc CreditCards before update

	db.Preload("CreditCards.Images").First(&updateUser, 1)
	fmt.Printf("updateUser after update: %+v\n", updateUser)
}
