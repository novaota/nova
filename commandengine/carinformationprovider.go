package main

import (
	"nova/car/models"
	"nova/rest"
	"nova/shared"
)

type CarInformationProvider interface {
	GetCurrentStatus() (models.CarInformationModel, error)
}

func NewLocalCarInformationProvider(certSettings shared.CertificateSettings, endpoint string) *localCarInformationProvider {
	return &localCarInformationProvider{
		Client: rest.NewClient(certSettings, endpoint),
	}
}

type localCarInformationProvider struct {
	*rest.Client
}

func (p *localCarInformationProvider) GetCurrentStatus() (models.CarInformationModel, error) {
	result := &models.CarInformationModel{}
	err := p.Get(result, "car/statusW")

	return *result, err
}
