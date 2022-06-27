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

func Stamps_by_user(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userId"))
	if responce, err := database.StampReadById(userid); err != nil {
		log.Println(err)
	} else {
		c.JSON(200, gin.H{"message": responce})
	}
}
