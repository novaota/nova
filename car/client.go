package main

import (
	"nova/car/models"
	"nova/rest"
	"nova/shared"
)

func NewCarClient(certSettings shared.CertificateSettings ,endpoint string) *carClient {
  return &carClient{
		Client: rest.NewClient(certSettings, endpoint),
	}
}

type carClient struct {
	*rest.Client
}

func (c *carClient) GetStatus() (*models.CarInformationModel, error) {
	model := &models.CarInformationModel{}
	err := c.Get(model,"/car/status")

	if err != nil {
		return nil, err
	}

	return model, nil
}
