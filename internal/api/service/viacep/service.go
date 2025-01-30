package viacep

import (
	"encoding/json"
	"go-kubernetes-poc/config"
	"io"
	"net/http"
	"strings"
)

type ViaCepServiceInterface interface {
	GetAddressByZipCode(zipCode string) (*ViaCepResponse, error)
}

type ViaCepService struct {
}

func NewViaCepService() *ViaCepService {
	return &ViaCepService{}
}

func (v *ViaCepService) GetAddressByZipCode(zipCode string) (*ViaCepResponse, error) {
	url := strings.Replace(config.Config.VIACEPURL, "$$cep", zipCode, 1)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	viacepResponse := &ViaCepResponse{}
	if err := json.Unmarshal(body, viacepResponse); err != nil {
		return nil, err
	}

	return viacepResponse, nil
}
