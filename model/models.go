package model

import (
	"time"
	"gorm.io/gorm"
)

type Stamps struct {
	gorm.Model
	UsersID 		int			`gorm:"not null" binding:"required"`
	Type			string		`gorm:"not null" binding:"required"`
	Stamp_datetime	time.Time	`gorm:"not null" binding:"required"`
}

type Users struct {
	ID           int			`gorm:"primaryKey"`
	FirstName string			`gorm:"not null" binding:"required"`
	LastName  string			`gorm:"not null" binding:"required"`
	Stamps  []Stamps			`gorm:"foreignKey:UsersID"`
}