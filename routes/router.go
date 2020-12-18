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
			user.POST("/getUserByID", controller.GetUserByID)
			user.POST("/register", controller.Register)
			user.POST("/login", controller.Login)
			user.POST("/logout", controller.Logout)
			user.POST("/getUserInfo", controller.GetUserInfo)
			user.POST("/updateUser", controller.UpdateUser)
			user.POST("/deleteUser", controller.DeleteUser)
			user.POST("/setAdmin", controller.SetUserAdmin)
		}

		role := api.Group("/role")
		{
			role.POST("/getAllRole", controller.GetAllRole)
		}

		userRole := api.Group("/userRole")
		{
			userRole.POST("/getUserRoles", controller.GetUserRoles)
			userRole.POST("/addUserRoles", controller.AddUserRoles)
			userRole.POST("/deleteUserRoles", controller.DeleteUserRoles)
		}
	}
	router.StaticFS("/public", http.Dir("./web"))
}