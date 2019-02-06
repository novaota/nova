package implementation

import (
	"nova/devicemanagement/datastorage"
	"nova/devicemanagement/datastorage/entities"
	"nova/devicemanagement/mappers"
	"nova/devicemanagement/models"
)

type deviceService struct {
	datastorage.RepositoryFactory
}

func NewDeviceService(factory datastorage.RepositoryFactory) *deviceService {
	return &deviceService{factory}
}

func (service *deviceService) GetAll() []models.CarModel {
	repo := service.CreateRepository()
	defer repo.Close()

	devices := &[]entities.Device{}
	repo.GetAll(devices)


	return mappers.CarMapper.MultipleMapFrom(*devices)
}

func (service *deviceService) GetByOwner(ownerModel models.OwnerModel) []models.CarModel {
	repo := service.CreateRepository()
	defer repo.Close()

	owner := mappers.OwnerMapper.MapTo(ownerModel)

	result := &[]entities.Device{}
	query := &entities.Device{Owner: owner}

	repo.GetWhere(result, query)

	return mappers.CarMapper.MultipleMapFrom(*result)
}

func (service *deviceService) Add(modelNr string, carModelModel *models.CarModel, ownerModel *models.OwnerModel) error {
	repo := service.CreateRepository()
	defer repo.Close()

		device := &entities.Device{
		ModelNr: modelNr,
		DeviceModelID: int(carModelModel.ID),
		OwnerID: int(ownerModel.ID),
	}

	err := repo.Add(&device)

	if err != nil {
		return err
	}

	return nil
}

func (service *deviceService) Update(carModel *models.CarModel) error {
	repo := service.CreateRepository()
	defer repo.Close()

	device := mappers.CarMapper.MapTo(*carModel)

	return repo.Update(&device)
}
