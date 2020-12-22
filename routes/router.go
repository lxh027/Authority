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
			role.POST("/getRoleByID", controller.GetRoleByID)
			role.POST("/addRole", controller.AddRole)
			role.POST("/deleteRole", controller.DeleteRole)
			role.POST("/updateRole", controller.UpdateRole)
		}

		userRole := api.Group("/userRole")
		{
			userRole.POST("/getUserRolesList", controller.GetUserRolesList)
			userRole.POST("/addUserRoles", controller.AddUserRoles)
			userRole.POST("/deleteUserRoles", controller.DeleteUserRoles)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/getAllAuth", controller.GetAllAuth)
			auth.POST("/getParentAuth", controller.GetParentAuth)
			auth.POST("/addAuth", controller.AddAuth)
			auth.POST("/deleteAuth", controller.DeleteAuth)
			auth.POST("/getAuthByID", controller.GetAuthByID)
			auth.POST("/updateAuth", controller.UpdateAuth)
		}
		roleAuth :=api.Group("roleAuth")
		{
			roleAuth.POST("/getRoleAuthsList", controller.GetRoleAuthsList)
			roleAuth.POST("/addRoleAuths", controller.AddRoleAuths)
		}
	}
	router.StaticFS("/public", http.Dir("./web"))
}