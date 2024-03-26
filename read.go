package main

import (
	"github.com/sunsunskibiz/gorm/model"
	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) (model.User, error) {
	var users model.User
	err := db.Model(&model.User{}).Preload("CreditCards.Images").Find(&users).Error
	return users, err
}

type Result struct {
	ID          string
	Name        string
	CreditCards []model.CreditCard `gorm:"foreignKey:UserID"`
}

func GetResult(db *gorm.DB) (Result, error) {
	var r Result
	err := db.Model(&model.User{}).
		Preload("CreditCards.Images").
		Find(&r).
		Where("name = ?", "MU").
		Error
	return r, err
}
