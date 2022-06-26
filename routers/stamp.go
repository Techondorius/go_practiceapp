package routers

import (
	"strconv"
	"time"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go_practiceapp/database"
	"go_practiceapp/model"
)

func User_in(c *gin.Context){
	userid, _ := strconv.Atoi(c.Param("userId"))

	form := model.Stamps{
		In_datetime: time.Now(),
		UsersID: userid,
	}

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		log.Println(err)
		return
	}
	if err := database.Stamp_create(&form); err != nil {
		c.JSON(400, gin.H{ "status": err.Error(), })
		log.Println(err)
	} else {
		c.JSON(200, gin.H{ "message": "list", })
	}
}

func User_up(c *gin.Context){
	userid, _ := strconv.Atoi(c.Param("userId"))
	up_datetime := time.Now()

	if err := database.Put_up_datetime(userid, up_datetime); err != nil {
		c.JSON(400, gin.H{ "status": err.Error(), })
		log.Println(err)
	} else {
		c.JSON(200, gin.H{ "message": "list", })
	}
}

func Stamp_delete(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("stampId"))
	if err := database.Delete_stamp(id); err != nil {
		c.JSON(400, gin.H{ "status": err.Error(), })
		log.Println(err)
	} else {
		c.JSON(200, gin.H{ "message": "list", })
	}
}