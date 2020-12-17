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
			user.POST("/login", controller.Login)
			user.POST("/logout", controller.Logout)
			user.POST("/getUserInfo", controller.GetUserInfo)
		}
	}
	router.StaticFS("/public", http.Dir("./web"))
}