package main

import (
	"go_practiceapp/database"
	"go_practiceapp/model"
)

func main() {
	db := database.Connection()

	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Stamps{})
}