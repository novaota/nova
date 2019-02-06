package services

import "nova/devicemanagement/models"

type OwnerService interface {
	GetAll() []models.OwnerModel
	Add(model models.OwnerModel) error
}