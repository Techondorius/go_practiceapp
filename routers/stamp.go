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
	var form model.Stamps

	form = model.Stamps{
		Stamp_datetime: time.Now(),
		UsersID: userid,
	}

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		log.Println(err)
		return
	}
	if err := database.Stamp_in(&form, "in"); err != nil {
		c.JSON(400, gin.H{ "message": "list", })
		log.Println(err)
	} else {
		c.JSON(200, gin.H{ "message": "list", })
	}
}

func User_up(c *gin.Context){
	userid, _ := strconv.Atoi(c.Param("userId"))
	var form model.Stamps

	form = model.Stamps{
		Stamp_datetime: time.Now(),
		UsersID: userid,
	}

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		log.Println(err)
		return
	}
	if err := database.Stamp_in(&form, "up"); err != nil {
		c.JSON(400, gin.H{ "message": "list", })
		log.Println(err)
	} else {
		c.JSON(200, gin.H{ "message": "list", })
	}
}