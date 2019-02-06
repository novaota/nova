package mappers

import (
	"nova/devicemanagement/datastorage/entities"
	"nova/devicemanagement/models"
)

var (
	OwnerMapper      ownerMapper =      ownerMapper{}
	UpdateTaskMapper updateTaskMapper = updateTaskMapper{}
	UpdateMapper     updateMapper =     updateMapper{}
	CarMapper        carMapper =        carMapper{}
	CarModelMapper   carModelMapper =   carModelMapper{}
)

//region car mapper

type carMapper struct{}

func (m *carMapper) MapTo(model models.CarModel) entities.Device {
	return entities.Device{
		OwnerID:           model.OwnerID,
		DeviceModelID:     model.DeviceModelID,
		ModelNr:           model.ModelNr,
		SoftwareVersion:   model.SoftwareVersion,
		CurrentBattery:    model.CurrentBattery,
		LastConnection:    model.LastConnection,
	}
}

func (m *carMapper) MapFrom(entity entities.Device) models.CarModel {
	return models.CarModel{
		OwnerID:         entity.OwnerID,
		DeviceModelID:   entity.DeviceModelID,
		ModelNr:         entity.ModelNr,
		SoftwareVersion: entity.SoftwareVersion,
		CurrentBattery:  entity.CurrentBattery,
		LastConnection:  entity.LastConnection,
	}
}

func (m *carMapper) MultipleMapFrom(entities []entities.Device) []models.CarModel {
	result := []models.CarModel{}

	for _, entity := range entities {
		model := m.MapFrom(entity)
		result = append(result, model)
	}

	return result
}

func (m *carMapper) MultipleMapTo(models []models.CarModel) []entities.Device {
	result := []entities.Device{}

	for _, model := range models {
		entity := m.MapTo(model)
		result = append(result, entity)
	}

	return result
}

//endregion

//region update mapper

type updateMapper struct{}

func (m *updateMapper) MapTo(model models.UpdateModel) entities.Update {
  result := entities.Update{
  	CarMustStand: model.CarMustStand,
		MinimalBattery: model.MinimalBattery,
		PayloadURL: model.PayloadURL,
		SuitableDeviceModelID: model.SuitableDeviceModelID,
		SuitableVersion: model.SuitableVersion,
  }

  result.ID = model.ID
  return result
}

func (m *updateMapper) MapFrom(entity entities.Update) models.UpdateModel {
	result := models.UpdateModel{
		SuitableVersion:       entity.SuitableVersion,
		SuitableDeviceModelID: entity.SuitableDeviceModelID,
		PayloadURL:            entity.PayloadURL,
		CarMustStand:          entity.CarMustStand,
		MinimalBattery:        entity.MinimalBattery,
	}

	result.ID = entity.ID
	return result
}

func (m *updateMapper) MultipleMapFrom(entities []entities.Update) []models.UpdateModel {
	result := []models.UpdateModel{}

	for _, entity := range entities {
		model := m.MapFrom(entity)
		result = append(result, model)
	}

	return result
}

func (m *updateMapper) MultipleMapTo(models []models.UpdateModel) []entities.Update {
	result := []entities.Update{}

	for _, model := range models {
		entity := m.MapTo(model)
		result = append(result, entity)
	}

	return result
}

//endregion

//region update task mapper

type updateTaskMapper struct {}

func (m updateTaskMapper) MapTo(model models.UpdateTaskModel) entities.UpdateTask {
	result := entities.UpdateTask{
		UpdateID:   model.UpdateID,
		DeviceID:   model.DeviceID,
		ExecutedAt: model.ExecutedAt,
		ReceivedAt: model.ReceivedAt,
		Status:     model.Status,
	}

	result.ID = model.ID
	return result
}

func (m *updateTaskMapper) MapFrom(entity entities.UpdateTask) models.UpdateTaskModel {
	result := models.UpdateTaskModel{
		UpdateID:              entity.UpdateID,
		DeviceID:              entity.DeviceID,
		MinimalBattery:        entity.Update.MinimalBattery,
		CarMustStand:          entity.Update.CarMustStand,
		PayloadURL:            entity.Update.PayloadURL,
		SuitableDeviceModelID: entity.Update.SuitableDeviceModelID,
		Status:                entity.Status,
		ExecutedAt:            entity.ExecutedAt,
		ReceivedAt:            entity.ReceivedAt,
	}

	result.ID = entity.ID
	return result
}

func (m *updateTaskMapper) MultipleMapFrom(entities []entities.UpdateTask) []models.UpdateTaskModel {
	result := []models.UpdateTaskModel{}

	for _, entity := range entities {
		model := m.MapFrom(entity)
		result = append(result, model)
	}

	return result
}

func (m *updateTaskMapper) MultipleMapTo(models []models.UpdateTaskModel) []entities.UpdateTask {
	result := []entities.UpdateTask{}

	for _, model := range models {
		entity := m.MapTo(model)
		result = append(result, entity)
	}

	return result
}

//endregion

//region owner mapper

type ownerMapper struct {}

func (m *ownerMapper) MapTo(model models.OwnerModel) entities.Owner {
	result := entities.Owner{
		Name:       model.Name,
		Zip:        model.Zip,
		Address:    model.Address,
		FirstName:  model.FirstName,
	}

	result.ID = model.ID
	return result
}

func (m *ownerMapper) MapFrom(entity entities.Owner) models.OwnerModel {
	result := models.OwnerModel{
		Name:       entity.Name,
		Zip:        entity.Zip,
		Address:    entity.Address,
		FirstName:  entity.FirstName,
	}

	result.ID = entity.ID
	return result
}

func (m *ownerMapper) MultipleMapFrom(entities []entities.Owner) []models.OwnerModel {
	result := []models.OwnerModel{}

	for _, entity := range entities {
		model := m.MapFrom(entity)
		result = append(result, model)
	}

	return result
}

func (m *ownerMapper) MultipleMapTo(models []models.OwnerModel) []entities.Owner {
	result := []entities.Owner{}

	for _, model := range models {
		entity := m.MapTo(model)
		result = append(result, entity)
	}

	return result
}

//endregion

//region user mapper

type userMapper struct {}

func (m *userMapper) MapTo(model models.UserModel) entities.User {
	result := entities.User{
		Username: model.Username,
		Password: []byte(model.Password),
		Role: model.Role,
	}

	result.ID = model.ID
	return result
}

func (m *userMapper) MapFrom(entity entities.User) models.UserModel {
  result := models.UserModel{
  	Role: entity.Role,
  	Username: entity.Username,
  	Password: "-----",
	}

  result.ID = entity.ID
  return result
}

func (m *userMapper) MultipleMapFrom(entities []entities.User) []models.UserModel {
	result := []models.UserModel{}

	for _, entity := range entities {
		model := m.MapFrom(entity)
		result = append(result, model)
	}

	return result
}

func (m *userMapper) MultipleMapTo(models []models.UserModel) []entities.User {
	result := []entities.User{}

	for _, model := range models {
		entity := m.MapTo(model)
		result = append(result, entity)
	}

	return result
}


//endregion

//region car model mapper

type carModelMapper struct {}

func (m *carModelMapper) MapTo(model models.CarModelModel) entities.DeviceModel {
	result := entities.DeviceModel{
		Name: model.Name,
	}

	result.ID = model.ID
	return result
}

func (m *carModelMapper) MapFrom(entity entities.DeviceModel) models.CarModelModel {
	result := models.CarModelModel{
		Name: entity.Name,
	}

	result.ID = entity.ID
	return result
}

func (m *carModelMapper) MultipleMapFrom(entities []entities.DeviceModel) []models.CarModelModel {
	result := []models.CarModelModel{}

	for _, entity := range entities {
		model := m.MapFrom(entity)
		result = append(result, model)
	}

	return result
}

func (m *carModelMapper) MultipleMapTo(models []models.CarModelModel) []entities.DeviceModel {
	result := []entities.DeviceModel{}

	for _, model := range models {
		entity := m.MapTo(model)
		result = append(result, entity)
	}

	return result
}

//endregion
