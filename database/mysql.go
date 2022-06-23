package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
    "go_practiceapp/model"
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

	if result.Error != nil {
		return result.Error
	}
	return nil
}