package model

type Users struct {
	ID           int	`gorm:"primaryKey"`
	FirstName string	`gorm:"not null" binding:"required"`
	LastName  string	`gorm:"not null" binding:"required"`
	Stamps  []Stamps	`gorm:"foreignKey:UsersID"`
}
