package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/sunsunskibiz/gorm/model"
)

var (
	one = "1"
	two = "2"
	three = "3"
	four = "4"
	five = "5"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sun?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Drop table
	db.Migrator().DropTable(&model.User{}, &model.CreditCard{}, &model.Image{})

	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.CreditCard{}, &model.Image{})

	// Create
	db.Create(&model.User{Name: "MU", ID: &one})
	db.Create(&model.CreditCard{Number: "111", UserID: &one, ID: &one})
	db.Create(&model.CreditCard{Number: "222", UserID: &one, ID: &two})
	db.Create(&model.Image{Url: "url-one-one", CreditCardID: "1", ID: "1"})
	db.Create(&model.Image{Url: "url-one-two", CreditCardID: "1", ID: "2"})

	// Read
	var user model.User
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
	var userWithAsso model.User
	db.Preload("CreditCards.Images").First(&userWithAsso, 1)
	fmt.Printf("userWithAsso before update: %+v\n", userWithAsso)

	userWithAsso.Name = "pikachu"
	userWithAsso.CreditCards = append(userWithAsso.CreditCards, model.CreditCard{Number: "333", UserID: &one, ID: &three})
	db.Updates(&userWithAsso)

	db.Preload("CreditCards.Images").First(&userWithAsso, 1)
	fmt.Printf("userWithAsso after update: %+v\n", userWithAsso)

	// Update user with partial field (just name)
	updateUser := model.User{
		ID: &one,
		Name: "charmander",
		CreditCards: []model.CreditCard{{Number: "444", UserID: &one, ID: &four}},
	}
	fmt.Printf("updateUser before update: %+v\n", updateUser)
	
	
	db.Debug().Omit("CreditCards").Updates(&updateUser)

	db.Preload("CreditCards.Images").First(&updateUser, 1)
	fmt.Printf("updateUser after update: %+v\n", updateUser)


	db.Debug().Model(&updateUser).Association("CreditCards").Replace([]model.CreditCard{{Number: "555", UserID: &one, ID: &five}})

	db.Preload("CreditCards.Images").First(&updateUser, 1)
	fmt.Printf("updateUser after update replace creditcards: %+v\n", updateUser)

}
