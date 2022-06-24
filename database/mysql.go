package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
    "go_practiceapp/model"
	"log"
	"errors"
)

func Connection() *gorm.DB {
	dsn := "root:@tcp(db:3306)/gin_app?charset=utf8&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db

}

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