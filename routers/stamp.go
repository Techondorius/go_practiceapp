package routers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"go_practiceapp/database"
	"go_practiceapp/model"
)

func User_in(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userId"))

	form := model.Stamps{
		In_datetime: time.Now(),
		UsersID:     userid,
	}

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Println(err)
		return
	}
	if user, err := database.CreateStamp(&form); err != nil {
		c.JSON(400, gin.H{"status": "Could not find user by id "})
		log.Println(err)
	} else {
		c.JSON(200, gin.H{
			"detail":  user,
			"message": "Stamped successfully",
		})
	}
}

func User_up(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userId"))
	up_datetime := time.Now()

	if err := database.StampPutUpTime(userid, up_datetime); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		log.Println(err)
	} else {
		c.JSON(200, gin.H{"message": "list"})
	}
}

func Stamp_delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("stampId"))
	if err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		log.Println(err)
	} else if err := database.DeleteStamp(id); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		log.Println(err)
	} else {
		c.JSON(200, gin.H{"message": "successfully up-ed"})
	}
}
