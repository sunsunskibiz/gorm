package model

import (
	"gorm.io/plugin/optimisticlock"
)

type User struct {
	ID          *string `gorm:"column:id;type:varbinary(30);primaryKey"`
	Name        string
	CreditCards []CreditCard `gorm:"foreignKey:UserID"`
	Version     optimisticlock.Version `gorm:"column:opt_lock;type:int;not null;default:0"`
}

type CreditCard struct {
	ID      *string `gorm:"column:id;type:varbinary(30);primaryKey"`
	Number  string
	UserID  *string
	Images  []Image 
	Version optimisticlock.Version
}

type Image struct {
	ID           string `gorm:"column:id;type:varbinary(30);primaryKey"`
	Url          string
	CreditCardID string
}
