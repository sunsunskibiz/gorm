package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
	Images []string
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sun?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Migrator().DropTable(&User{}, &CreditCard{})

	// Migrate the schema
	db.AutoMigrate(&User{}, &CreditCard{})

	// Create
	db.Create(&User{Name: "MU"})
	db.Create(&CreditCard{Number: "111", UserID: 1})
	db.Create(&CreditCard{Number: "222", UserID: 1})

	// Read
	var user User
	db.First(&user, 1)
	fmt.Printf("User: %+v\n", user)

	users, err := GetAll(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Users: %+v\n", users)
}

func GetAll(db *gorm.DB) (User, error) {
    var users User
    err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
    return users, err
}
