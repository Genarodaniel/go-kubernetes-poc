package address

import (
	"go-kubernetes-poc/internal/api/service/viacep"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	viacep := viacep.NewViaCepService()
	service := NewAddressService(viacep)
	handler := NewAddressHandler(service)

	g.GET("/:zipcode", handler.HandleGetAddressByZipCode)
}
