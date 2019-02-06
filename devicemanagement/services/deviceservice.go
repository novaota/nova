package services

import (
	"nova/devicemanagement/models"
)

type DeviceService interface {
	Update(device *models.CarModel) error
	GetAll() []models.CarModel
	Add(modelNr string, model *models.CarModel, owner *models.OwnerModel) error
	GetByOwner(owner models.OwnerModel) []models.CarModel
}
