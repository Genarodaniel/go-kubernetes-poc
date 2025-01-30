package address

import (
	"go-kubernetes-poc/internal/api/service/viacep"
)

type AddressServiceInterface interface {
	GetAddressByZipCode(zipCode string) (*GetAddressByZipCodeResponse, error)
}

type AddressService struct {
	ViaCepService viacep.ViaCepServiceInterface
}

func NewAddressService(viaCepService viacep.ViaCepServiceInterface) *AddressService {
	return &AddressService{
		ViaCepService: viaCepService,
	}
}

func (a *AddressService) GetAddressByZipCode(zipCode string) (*GetAddressByZipCodeResponse, error) {
	viacepResponse, err := a.ViaCepService.GetAddressByZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	return &GetAddressByZipCodeResponse{
		ZipCode:      viacepResponse.Cep,
		Complement:   viacepResponse.Complemento,
		Street:       viacepResponse.Logradouro,
		Neighborhood: viacepResponse.Bairro,
		City:         viacepResponse.Localidade,
		StateCode:    viacepResponse.Uf,
		State:        viacepResponse.Estado,
		Region:       viacepResponse.Regiao,
	}, nil
}
