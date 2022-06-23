package routers

import (
	"go_practiceapp/model"
	"github.com/gin-gonic/gin"
	"go_practiceapp/database"
	"strconv"
)

func User_in(c *gin.Context){
	db := database.Connection()
    // defer db.Close()

	userid, _ := strconv.Atoi(c.Param("userId"))
	db.Create(&model.Stamps{UsersID: userid, Type: "in"})
	c.JSON(200, gin.H{ "message": "list", })

}