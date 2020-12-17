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
			user.POST("/getAllUser", controller.GetAllUser)
			user.POST("/register", controller.Register)
			user.POST("/login", controller.Login)
			user.POST("/logout", controller.Logout)
			user.POST("/getUserInfo", controller.GetUserInfo)
			user.POST("/updateUser", controller.UpdateUser)
			user.POST("/deleteUser", controller.DeleteUser)
		}
	}
	router.StaticFS("/public", http.Dir("./web"))
}