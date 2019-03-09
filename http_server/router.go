package http_server

import (
	"github.com/eager7/eos_docker_tools/http_server/handle"
	"github.com/gin-gonic/gin"
)

func InitializeGin() error {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	initRouter(engine)

	return engine.Run("localhost:13141")
}

func initRouter(router *gin.Engine) {
	router.GET("/ping", handle.Ping)
	router.GET("/version", handle.Version)

	router.POST("/contracts", handle.ContractsHandle)
}
