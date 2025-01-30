package address_test

import (
	"errors"
	"go-kubernetes-poc/internal/api/address"
	"go-kubernetes-poc/internal/api/service/viacep"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetAddressByZipCode(t *testing.T) {
	gin.SetMode(gin.TestMode)
	address.Router(&gin.Default().RouterGroup)
	path := "/address/v1/"

	t.Run("Should return error when payload is empty", func(t *testing.T) {
		viacepSpy := &viacep.ViaCepSpy{
			GetAddressByZipCodeResponse: nil,
			GetAddressByZipCodeError:    nil,
		}

		addressService := address.NewAddressService(viacepSpy)

		addressHandler := address.NewAddressHandler(addressService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "zipcode", Value: ""})
		addressHandler.HandleGetAddressByZipCode(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Should return error when zipcode is not an integer", func(t *testing.T) {
		viacepSpy := &viacep.ViaCepSpy{
			GetAddressByZipCodeResponse: nil,
			GetAddressByZipCodeError:    nil,
		}

		addressService := address.NewAddressService(viacepSpy)

		addressHandler := address.NewAddressHandler(addressService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "zipcode", Value: "123456ab"})
		addressHandler.HandleGetAddressByZipCode(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, string(response), "zipcode must be an integer")
	})

	t.Run("Should return an error when calling viacep", func(t *testing.T) {
		serviceError := "Error to call viacep"
		viacepSpy := &viacep.ViaCepSpy{
			GetAddressByZipCodeResponse: nil,
			GetAddressByZipCodeError:    errors.New(serviceError),
		}

		addressService := address.NewAddressService(viacepSpy)

		addressHandler := address.NewAddressHandler(addressService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "zipcode", Value: "12312312"})
		addressHandler.HandleGetAddressByZipCode(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, string(response), serviceError)
	})

	t.Run("Should return success", func(t *testing.T) {
		viacepSpy := &viacep.ViaCepSpy{
			GetAddressByZipCodeResponse: &viacep.ViaCepResponse{},
			GetAddressByZipCodeError:    nil,
		}

		addressService := address.NewAddressService(viacepSpy)

		addressHandler := address.NewAddressHandler(addressService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "zipcode", Value: "01530000"})
		addressHandler.HandleGetAddressByZipCode(ctx)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}
