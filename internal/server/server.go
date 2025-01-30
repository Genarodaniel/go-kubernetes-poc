package server

import (
	"go-kubernetes-poc/internal/api"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	//update with config env value
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	router.Use(gin.Recovery())

	api.Router(router)

	return router
}
