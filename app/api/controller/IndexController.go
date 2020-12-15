package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Index(c *gin.Context)  {

	c.JSON(http.StatusOK, gin.H{
		"msg": "Car Management",
	})
}