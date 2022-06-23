package model

// "gorm.io/gorm"

type Users struct {
	ID        int
	FirstName string
	LastName  string
	Stamps    Stamps
}
