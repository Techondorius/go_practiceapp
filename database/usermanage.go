package database

import (
    "go_practiceapp/model"
	"log"
	"errors"
)

func Create_User(User *model.Users) error {
	db := Connection()
	user := model.Users{FirstName: User.FirstName, LastName: User.LastName}
	result := db.Create(&user)
	log.Println(result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Update_User(User *model.Users) error {
	db := Connection()
	var user model.Users
	result := db.Model(&user).Where("id = ?", User.ID).Updates(User)
	log.Println(result.RowsAffected)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}