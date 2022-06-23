package routers

import "github.com/gin-gonic/gin"

func NewUser(c *gin.Context){
	c.JSON(200, gin.H{ "message": "list", })
}

func EditUser(c *gin.Context){

}

func DeleteUser(c *gin.Context){

}