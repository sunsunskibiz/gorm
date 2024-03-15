package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID      string
	Name        string
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	CreditCardID string
	Number       string
	UserID       string
	Images       []Image
}

type Image struct {
	gorm.Model
	Url          string
	CreditCardID string
}