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

func StampIn(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userId"))
	user, err2 := database.ReadUserByID(userid)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err2.Error()})
		log.Println(err2)
		return
	}

	f := model.Stamps{
		In_datetime: time.Now(),
		UsersID:     userid,
		Hourly_wage: user.Hourly_wage,
	}
	log.Println(f.In_datetime)
	if err := c.Bind(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Println(err)
		return
	}
	if user, err := database.CreateStamp(f); err != nil {
		c.JSON(400, gin.H{"status": "Could not find user by id "})
		log.Println(err)
	} else {
		c.JSON(200, gin.H{
			"detail":  map[string]any{
				"ID": user.ID,
				"UsersID": user.UsersID,
				"In_datetime": user.In_datetime.Format("2006/01/02 03:04"),
				"Hourly_wage": user.Hourly_wage,
			},
			"message": "Stamped successfully",
		})
	}
}

func StampUp(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userId"))
	up_datetime := time.Now()

	if _, err := database.UpdateStampUpTime(userid, up_datetime); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		log.Println(err)
	} else {
		c.JSON(200, gin.H{"message": "list"})
	}
}

func StampUpdate(c *gin.Context) {
	stampid, err := strconv.Atoi(c.Param("stampId"))
	if err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		log.Println(err)
		return
	}

	var times model.StampsDatetime
	if err2 := c.Bind(&times); err2 != nil{
		c.JSON(400, gin.H{"status": err.Error()})
		log.Println(err)
		return
	}

	indatetime, _ := time.Parse("2006/01/02 15:04:05 (MST)", times.In_datetime + " (JST)")
	updatetime, _ := time.Parse("2006/01/02 15:04:05 (MST)", times.Up_datetime + " (JST)")

	if r, err := database.UpdateStampTimestamp(stampid, indatetime, updatetime); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": r.OnlyDatetimes()})
	}
}

func StampDelete(c *gin.Context) {
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
