package main

import (
	"github.com/pkg/errors"
	"nova/communication"
)

type updateValidator struct {
	provider CarInformationProvider
}

func NewUpdateValidator(provider CarInformationProvider) *updateValidator {
  return &updateValidator{
  	provider: provider,
	}
}

func (val *updateValidator) ValidateRequirements(requirements communication.UpdateRequirements) (bool, error) {
	currentStatus, err := val.provider.GetCurrentStatus()

	if err != nil {
		return false, err
	}

	if currentStatus.BatteryLevel < requirements.MinimalBattery {
		return false, errors.New("Battery is too low")
	}

	if currentStatus.IsDriving && requirements.CarMustStand {
		return false, errors.New("Update can only be applied, when car is standing")
	}

	if currentStatus.NetworkReception < requirements.MinimalNetworkReception {
		return false, errors.New("Network reception is too low")
	}

	// yaay success
	return true, nil
}