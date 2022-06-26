package model

import (
	"time"
	// "gorm.io/gorm"
)

type Stamps struct {
	ID				int			`gorm:"primaryKey"`
	UsersID 		int			`gorm:"not null;" binding:"required"`
	In_datetime		time.Time
	Up_datetime 	*time.Time
}

type Users struct {
	ID           int			`gorm:"primaryKey"`
	FirstName string			`gorm:"not null" binding:"required"`
	LastName  string			`gorm:"not null" binding:"required"`
	Stamps  []Stamps			`gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Users_noStamps struct {
	ID           int
	FirstName string
	LastName  string
}