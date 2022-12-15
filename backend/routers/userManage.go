package routers

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"

	"go_practiceapp/database"
	"go_practiceapp/model"
)

func NewUser(c *gin.Context) {
	var form model.Users
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		log.Println(err)
		return
	}
	if responce, err := database.CreateUser(&form); err != nil {
		c.JSON(400, nil)
		return
		} else {
		c.JSON(200, gin.H{
			"detail": responce.DropStamps(),
			"message": "Created",
		})
		return
	}
}

func ShowUser(c *gin.Context) {
	if responce, err := database.ReadUser(); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"detail":  "something error has occured",
		})
	} else {
		var user_list []any
		for _, responce_i := range responce {
			user_list = append(user_list, responce_i.DropStamps())
		}

		c.JSON(200, gin.H{
			"detail": user_list,
		})
	}
}

func EditUser(c *gin.Context) {
	var form model.Users
	_ = c.Bind(&form)
	form.ID, _ = strconv.Atoi(c.Param("userId"))

	u, _ := database.ReadUserByID(form.ID)

	_ = mergo.Merge(&form, u)
	if user, err := database.UpdateUser(&form); err != nil {
		c.JSON(400, nil)
	} else {
		c.JSON(200, gin.H{
			"detail": user.DropStamps(),
			"message": "Updated",
		})
	}
}

func DeleteUser(c *gin.Context) {
	var form model.Users
	form.ID, _ = strconv.Atoi(c.Param("userId"))
	if err := database.DeleteUser(&form); err != nil {
		c.JSON(400, nil)
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Deleted",
		})
		c.Abort()
		return
	}
}
