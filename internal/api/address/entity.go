package address

import "errors"

type GetAddressByZipCodeResponse struct {
	ZipCode      string `json:"zip_code"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
	Unit         string `json:"unit"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	StateCode    string `json:"state_code"`
	State        string `json:"state"`
	Region       string `json:"region"`
}

type GetAddressByZipCodeRequest struct {
	ZipCode string `json:"zip_code"`
}

func (request *GetAddressByZipCodeRequest) Validate() error {
	if request.ZipCode == "" {
		return errors.New("zipcode must not be empty")
	}

	if len(request.ZipCode) != 8 {
		return errors.New("zipcode must have 8 characters")
	}

	return nil
}
