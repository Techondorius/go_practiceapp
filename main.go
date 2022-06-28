package main

import (
	"github.com/gin-gonic/gin"

	"go_practiceapp/routers"
)

func main() {

	r := gin.Default()

	userManage := r.Group("/userManage")
	{
		userManage.POST("/newUser", routers.NewUser)
		userManage.GET("/getUser", routers.ShowUser)
		userManage.PUT("/editUser/:userId", routers.EditUser)
		userManage.DELETE("/deleteUser/:userId", routers.DeleteUser)
	}

	stamp := r.Group("/stamp/:userId")
	{
		stamp.POST("/in", routers.StampIn)
		stamp.PUT("/up", routers.StampUp)
	}
	list := r.Group("/list")
	{
		list.GET("/byUser/:userId", routers.StampGetByUser)
	}

	edit := r.Group("/edit")
	{
		edit.PUT("/:stampId", routers.StampUpdate)
		edit.DELETE("/:stampId", routers.StampDelete)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})

	r.Run()
}