package routers

import (
	"strconv"
	"math"
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"

	"go_practiceapp/database"
	// "go_practiceapp/model"
)

func StampGetByUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userId"))
	stamps, err := database.ReadStampByUserId(userid)
	if err != nil {
		log.Println(err)
	} else {
		var fee float64 = 0
		var s []interface{}
		for _, res := range stamps {
			t := res.Up_datetime.Sub(res.In_datetime)
			fee += float64(res.Hourly_wage) * t.Minutes() / 60
			s = append(s, res.OnlyDatetimes())
		}
		c.JSON(200, gin.H{
			"stamps": s,
			"fee": math.Round(fee),
		})
	}
}
