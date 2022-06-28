package routers

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go_practiceapp/database"
	"go_practiceapp/model"
)

func NewUser(c *gin.Context) {
	var form model.Users
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"detail": "Request has bad value",
		})
		log.Println(err)
		return
	}
	log.Println(form.FirstName)
	if responce, err := database.CreateUser(&form); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"detail":  "something error has occured",
		})
	} else {
		c.JSON(200, gin.H{
			"detail": map[string]any{
				"ID":        responce.ID,
				"FirstName": responce.FirstName,
				"LastName":  responce.LastName,
			},
			"message": "Created",
		})
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
	err := c.Bind(&form)
	var err2 error
	form.ID, err2 = strconv.Atoi(c.Param("userId"))

	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"detail": "Request has bad value",
		})
		log.Println(err, err2)
		return
	}

	if err := database.UpdateUser(&form); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"detail":  "no records edited",
		})
	} else {
		c.JSON(200, gin.H{
			"detail": map[string]any{
				"ID":        form.ID,
				"FirstName": form.FirstName,
				"LastName":  form.LastName,
			},
			"message": "Updated",
		})
	}
}

func DeleteUser(c *gin.Context) {
	var form model.Users
	c.Bind(&form)
	form.ID, _ = strconv.Atoi(c.Param("userId"))
	if err := database.DeleteUser(&form); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"detail":  "no records edited",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Deleted",
		})
		c.Abort()
		return
	}
}
