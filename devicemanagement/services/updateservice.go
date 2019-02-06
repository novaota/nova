package services

import (
	"nova/devicemanagement/datastorage/entities"
	"nova/devicemanagement/models"
)

type UpdateService interface {
	GetUpdateById(id uint) models.UpdateModel
	GetUpdateTaskById(id uint) models.UpdateTaskModel
	GetAllUpdates() []models.UpdateModel
	GetAllUpdateTasks() []models.UpdateTaskModel
	CreateUpdate(update models.UpdateModel) error
	ChangeUpdate(update models.UpdateModel) error
	UpdateDevice(model models.UpdateTaskModel) error
  // GetUpdateStatus(device *models.CarModel)
	GetOutstandingTasks() []models.UpdateTaskModel
	SetUpdateStatus(updateTaskID uint, status entities.CommandStatus)
}
