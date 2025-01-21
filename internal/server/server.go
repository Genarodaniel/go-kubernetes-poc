package server

import (
	"address-crud-1/config"
	"address-crud-1/internal/api"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(config.Config.GinMode)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	router.Use(gin.Recovery())

	api.Router(router)

	return router
}
