package services

import "nova/devicemanagement/models"

type DeviceModelService interface {
	GetAll() []models.CarModelModel
	Add(model models.CarModelModel) error
}