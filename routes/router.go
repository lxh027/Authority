package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine)  {

	// api
	router.StaticFS("/public", http.Dir("./web"))
}