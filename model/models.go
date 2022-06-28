package model

import (
	"time"
)

type Stamps struct {
	ID				int			`gorm:"primaryKey"`
	UsersID 		int			`gorm:"not null;" binding:"required"`
	In_datetime		time.Time
	Up_datetime 	*time.Time
	Hourly_wage		int			`gorm:"not null;" binding:"required"`
}

type Users struct {
	ID          int			`gorm:"primaryKey"`
	FirstName	string		`gorm:"not null" binding:"required"`
	LastName	string		`gorm:"not null" binding:"required"`
	Hourly_wage int			`gorm:"not null;default:1000"`
	Stamps  	[]Stamps	`gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


func (user Users) DropStamps() map[string]any {
	return map[string]any{
		"ID" : user.ID,
		"FirstName" : user.FirstName,
		"LastName" : user.LastName,
		"Hourly_wage" : user.Hourly_wage,
	}
}