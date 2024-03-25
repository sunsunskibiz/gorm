package main

import "gorm.io/gorm"

func GetAll(db *gorm.DB) (User, error) {
	var users User
	err := db.Model(&User{}).Preload("CreditCards.Images").Find(&users).Error
	return users, err
}

type Result struct {
	ID      string
	Name        string
	CreditCards []CreditCard `gorm:"foreignKey:UserID"`
}

func GetResult(db *gorm.DB) (Result, error) {
	var r Result
	err := db.Model(&User{}).
		Preload("CreditCards.Images").
		Find(&r).
		Where("name = ?", "MU").
		Error
	return r, err
}