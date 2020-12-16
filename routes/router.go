package routes

import (
	"Authority/app/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine)  {
	// api

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", controller.Register)
		}
	}
	router.StaticFS("/public", http.Dir("./web"))
}