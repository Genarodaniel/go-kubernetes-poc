package address_test

import (
	"go-kubernetes-poc/internal/api/address"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return an error when zipcode is empty", func(t *testing.T) {
		request := address.GetAddressByZipCodeRequest{}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "zipcode must not be empty")
	})

	t.Run("Should return an error when zipcode don't have 8 characters", func(t *testing.T) {
		request := address.GetAddressByZipCodeRequest{
			ZipCode: "123",
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "zipcode must have 8 characters")
	})

	t.Run("Should return success when the zipcode is valid", func(t *testing.T) {
		request := address.GetAddressByZipCodeRequest{
			ZipCode: "01530000",
		}
		err := request.Validate()

		assert.Nil(t, err)
	})

}
