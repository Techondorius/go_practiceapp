package routers

import (
	"go_practiceapp/model"
	"github.com/gin-gonic/gin"
	"go_practiceapp/database"
)

func NewUser(c *gin.Context){
	db := database.Connection()
    // defer db.Close()
	db.Create(&model.Users{FirstName: "Kyosuke", LastName: "Fujita"})
	c.JSON(200, gin.H{ "message": "list", })
}

func EditUser(c *gin.Context){

}

func DeleteUser(c *gin.Context){

}