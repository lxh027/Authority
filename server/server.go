package server

import (
	"Authority/config"
	"Authority/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Run(httpServer *gin.Engine)  {
	serverConfig  := config.GetServerConfig()
	sessionConfig := config.GetSessionConfig()
	// 运行模式
	gin.SetMode(serverConfig["mode"].(string))
	httpServer = gin.Default()

	// 创建session存储引擎
	sessionStore := cookie.NewStore([]byte(sessionConfig["key"].(string)))
	sessionStore.Options(sessions.Options{
		MaxAge: int(sessionConfig["time"].(int)*60000),
	})
	//设置session中间件
	httpServer.Use(sessions.Sessions(sessionConfig["name"].(string), sessionStore))

	gin.DisableConsoleColor()
	// 生成日志
	logFile, _:= os.Create(config.GetLogPath())
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	// 设置日志格式
	httpServer.Use(gin.LoggerWithFormatter(config.GetLogFormat))
	httpServer.Use(gin.Recovery())

	// 注册路由
	routes.Routes(httpServer)

	serverError := httpServer.Run(serverConfig["host"].(string)+":"+serverConfig["port"].(string))

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}
}