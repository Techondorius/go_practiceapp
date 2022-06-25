package database

import (
	"log"
	"errors"

	// "gorm.io/gorm"

	"go_practiceapp/model"
)

func Create_User(User *model.Users) (model.Users, error) {
	db := Connection()
	user := model.Users{FirstName: User.FirstName, LastName: User.LastName}
	result := db.Create(&user)
	log.Println(user.ID)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
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

func Delete_User(User *model.Users) error{
	db := Connection()
	var user model.Users
	// var stamp model.Stamps

	// result2 := db.Model(&stamp).Where("users_id = ?", User.ID).Unscoped().Delete(&stamp)
	// log.Println("result2", result2)
	result := db.Model(&user).Where("id = ?", User.ID).Delete(User)
	log.Println(result)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}