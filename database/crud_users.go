package database

import (
	"errors"

	// "gorm.io/gorm"

	"go_practiceapp/model"
)

func CreateUser(User *model.Users) (model.Users, error) {
	db := Connection()
	user := model.Users{FirstName: User.FirstName, LastName: User.LastName}
	result := db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func ReadUser() ([]model.Users, error) {
	var users []model.Users
	db := Connection()
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return users, result.Error
	}
}

func UpdateUser(User *model.Users) error {
	db := Connection()
	var user model.Users
	result := db.Model(&user).Where("id = ?", User.ID).Updates(User)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}

func DeleteUser(User *model.Users) error {
	db := Connection()
	var user model.Users
	result := db.Model(&user).Where("id = ?", User.ID).Delete(User)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}
