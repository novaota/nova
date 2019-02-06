package implementation

import (
	"nova/devicemanagement/datastorage"
	"nova/devicemanagement/datastorage/entities"
	"nova/devicemanagement/mappers"
	"nova/devicemanagement/models"
)

func NewDeviceModelService(factory datastorage.RepositoryFactory) *deviceModelService  {
	return &deviceModelService{factory}
}

type deviceModelService struct {
	datastorage.RepositoryFactory
}

func (s *deviceModelService) GetAll() []models.CarModelModel {
	repo := s.CreateRepository()
	defer repo.Close()

	result := []entities.DeviceModel{}
	repo.GetAll(result)

	return mappers.CarModelMapper.MultipleMapFrom(result)
}

func (s *deviceModelService) Add(model models.CarModelModel) error {
	repo := s.CreateRepository()
	defer repo.Close()

	entity := mappers.CarModelMapper.MapTo(model)

	return repo.Add(entity)
	}