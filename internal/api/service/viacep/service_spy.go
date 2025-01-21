package viacep

type ViaCepSpy struct {
	GetAddressByZipCodeResponse *ViaCepResponse
	GetAddressByZipCodeError    error
}

func (v *ViaCepSpy) GetAddressByZipCode(zipCode string) (*ViaCepResponse, error) {
	return v.GetAddressByZipCodeResponse, v.GetAddressByZipCodeError
}
