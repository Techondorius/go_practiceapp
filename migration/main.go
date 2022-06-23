package main

import (
	"go_practiceapp/database"
	"go_practiceapp/model"
)

func main() {
    db := database.Connection()
    // defer db.Close()

    db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Stamps{})
}