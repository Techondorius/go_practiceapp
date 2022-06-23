package model

import (
	"gorm.io/gorm"
)

type Stamps struct {
	gorm.Model
	UsersID int
	Type    string
}
