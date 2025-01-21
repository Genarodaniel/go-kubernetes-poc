package address

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressHandlerInterface interface {
	HandleGetAddressByZipCode(c *gin.Context)
}

type AddressHandler struct {
	AddressService AddressServiceInterface
}

func NewAddressHandler(addressService AddressServiceInterface) *AddressHandler {
	return &AddressHandler{
		AddressService: addressService,
	}
}

func (h *AddressHandler) HandleGetAddressByZipCode(ctx *gin.Context) {
	request := &GetAddressByZipCodeRequest{}
	request.ZipCode = ctx.Param("zipcode")
	if err := request.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	_, err := strconv.Atoi(request.ZipCode)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": fmt.Sprintf("zipcode must be an integer %s ", err.Error()),
		})
		return
	}

	response, err := h.AddressService.GetAddressByZipCode(request.ZipCode)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
