package routers

import (
	"strconv"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go_practiceapp/database"
	"go_practiceapp/model"
	// "encode/json"
)

func NewUser(c *gin.Context){
	var form model.Users
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"detail": "Request has bad value",
		})
		log.Println(err)
		return
	}
	if responce, err := database.Create_User(&form); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"detail": "something error has occured",
		})
	} else {
		c.JSON(200, gin.H{
			"detail": map[string]any{
				"ID": responce.ID,
				"FirstName": responce.FirstName,
				"LastName": responce.LastName,
			},
			"message": "Created",
		})
	}
}

func ShowUser(c *gin.Context){
	if responce, err :=database.Read_User(); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"detail": "something error has occured",
		})
	} else {
		var user_list []model.Users_noStamps
		for _, responce_i := range responce{
			user_list = append(user_list, model.Users_noStamps{
				ID : responce_i.ID,
				FirstName : responce_i.FirstName,
				LastName : responce_i.LastName,
			})
		}

		c.JSON(200, gin.H{
			"detail": user_list,
		})
	}
}

func EditUser(c *gin.Context){
	var form model.Users
	err := c.Bind(&form)
	var err2 error
	form.ID, err2 = strconv.Atoi(c.Param("userId"))
	log.Println(form)

	if err != nil || err2 != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"detail": "Request has bad value",
		})
		log.Println(err, err2)
		return
	}

	if err := database.Update_User(&form); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"detail": "no records edited",
		})
	} else {
		c.JSON(200, gin.H{
			"detail": map[string]any{
				"ID": form.ID,
				"FirstName": form.FirstName,
				"LastName": form.LastName,
			},
			"message": "Updated",
		})
	}
}

func DeleteUser(c *gin.Context){
	var form model.Users
	err := c.Bind(&form)
	var err2 error
	form.ID, err2 = strconv.Atoi(c.Param("userId"))
	log.Println(form)

	if err2 != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"detail": "Request has bad value",
		})
		log.Println(err, err2)
		return
	}

	if err := database.Delete_User(&form); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"detail": "no records edited",
		})
	} else {
		c.JSON(200, gin.H{
			"detail": map[string]any{
				"ID": form.ID,
			},
			"message": "Deleted",
		})
	}
}