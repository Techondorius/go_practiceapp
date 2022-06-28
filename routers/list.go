package routers

import (
	"strconv"
	// "time"
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"

	"go_practiceapp/database"
	// "go_practiceapp/model"
)

func StampGetByUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userId"))
	stamps, err := database.ReadStampById(userid)
	user, err2 := database.ReadUserByID(userid)
	if err != nil || err2 != nil {
		log.Println(err, err2)
	} else {
		fee := 0
		for _, res := range stamps {
			t := res.Up_datetime.Sub(res.In_datetime)
			fee += int(t.Minutes())
			log.Println(user)
		}
		c.JSON(200, gin.H{"message": stamps})
	}
}
