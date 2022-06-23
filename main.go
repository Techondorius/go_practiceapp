package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"go_practiceapp/routers"
	// "go_practiceapp/database"
)

func main(){

	r := gin.Default()

	userManage := r.Group("/userManage")
	{
		userManage.POST("/newUser", routers.NewUser)
		userManage.PUT("/editUser/:userId", routers.EditUser)
		userManage.DELETE("/deleteUser/:userId", routers.DeleteUser)
	}

	stamp := r.Group("/stamp/:userId")
	{
		stamp.POST("/in", routers.User_in)
		stamp.POST("/up", routers.User_in)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "good", })
	})

	r.Run()
}