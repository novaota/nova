package implementation

import (
	"nova/devicemanagement/datastorage"
	"nova/devicemanagement/datastorage/entities"
	"nova/devicemanagement/mappers"
	"nova/devicemanagement/models"
)

type ownerService struct {
	datastorage.RepositoryFactory
}

func NewOwnerService(factory datastorage.RepositoryFactory) *ownerService {
	return &ownerService{factory}
}


func (s *ownerService) GetAll() []models.OwnerModel {
	repo := s.CreateRepository()
	defer repo.Close()

	result := &[]entities.Owner{}

	repo.GetAll(result)

	return mappers.OwnerMapper.MultipleMapFrom(*result)
}

func (s *ownerService) Add(model models.OwnerModel) error {
	repo := s.CreateRepository()
	defer repo.Close()

	entity := mappers.OwnerMapper.MapTo(model)

	return repo.Add(entity)
}



