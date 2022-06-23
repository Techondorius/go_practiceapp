package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"go_practiceapp/routers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
    gorm.Model
    Code  string
    Price uint
}

func main(){

	db, err := gorm.Open("mysql", "root:@tcp(db:3306)/gin_app?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&Product{})

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