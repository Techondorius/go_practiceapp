package routers

import (
	"strconv"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go_practiceapp/database"
	"go_practiceapp/model"
)

func NewUser(c *gin.Context){
	var form model.Users
    if err := c.Bind(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		log.Println(err)
        return
    }
	if err := database.Create_User(&form); err != nil {
		c.JSON(400, gin.H{ "message": "Bad request", })
	} else {
		c.JSON(200, gin.H{ "message": "User Created", })
	}
}

func EditUser(c *gin.Context){
	var form model.Users
	err := c.Bind(&form)
	var err2 error
	form.ID, err2 = strconv.Atoi(c.Param("userId"))
	log.Println(form)

    if err != nil || err2 != nil{
        c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		log.Println(err, err2)
        return
    }

	if err := database.Update_User(&form); err != nil {
		c.JSON(400, gin.H{ "message": "Bad request", })
	} else {
		c.JSON(200, gin.H{ "message": "User Created", })
	}
}

func DeleteUser(c *gin.Context){

}