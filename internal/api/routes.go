package api

import (
	"go-kubernetes-poc/internal/api/healthcheck"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	v1 := e.Group("/v1")

	// addressGroup := v1.Group("/address")
	healthCheckGroup := v1.Group("/healthcheck")

	healthcheck.Router(healthCheckGroup)
	// address.Router(addressGroup)
}
