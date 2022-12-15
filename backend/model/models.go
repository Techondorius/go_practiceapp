package model

import (
	"time"
)

type Stamps struct {
	ID          int `gorm:"primaryKey"`
	UsersID     int `gorm:"not null;"`
	In_datetime time.Time
	Up_datetime *time.Time
	Hourly_wage int `gorm:"not null;"`
}

type Users struct {
	ID          int      `gorm:"primaryKey"`
	FirstName   string   `gorm:"not null"`
	LastName    string   `gorm:"not null"`
	Hourly_wage int      `gorm:"not null;default:1000"`
	Stamps      []Stamps `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StampsDatetime struct {
	In_datetime string
	Up_datetime string
}

func (user Users) DropStamps() map[string]any {
	return map[string]any{
		"ID":          user.ID,
		"FirstName":   user.FirstName,
		"LastName":    user.LastName,
		"Hourly_wage": user.Hourly_wage,
	}
}

func (stamp Stamps) OnlyDatetimes() map[string]any {
	return map[string]any{
		"ID":          stamp.ID,
		"In_datetime": stamp.In_datetime.Format("2006/01/02 03:04"),
		"Up_datetime": stamp.Up_datetime.Format("2006/01/02 03:04"),
	}
}
