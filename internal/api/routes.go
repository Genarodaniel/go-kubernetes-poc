package api

import (
	"address-crud-1/internal/api/address"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	v1 := e.Group("/v1")

	addressGroup := v1.Group("/address")
	address.Router(addressGroup)
}
