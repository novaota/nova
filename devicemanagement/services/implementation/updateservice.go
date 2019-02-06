package implementation

import (
	"nova/devicemanagement/datastorage"
	"nova/devicemanagement/datastorage/entities"
	"nova/devicemanagement/mappers"
	"nova/devicemanagement/models"
)

type updateService struct {
	datastorage.RepositoryFactory
}

func NewUpdateService(factory datastorage.RepositoryFactory) *updateService {
	return &updateService{factory}
}

func (service *updateService) CreateUpdate(updateModel models.UpdateModel) error {
	repo := service.CreateRepository()
	defer repo.Close()

	update := mappers.UpdateMapper.MapTo(updateModel)

	return repo.Add(update)
}

func (service *updateService) ChangeUpdate(updateModel models.UpdateModel) error {
	repo := service.CreateRepository()
	defer repo.Close()

	update := mappers.UpdateMapper.MapTo(updateModel)

	return repo.Update(&update)
}


func (service *updateService) UpdateDevice(updateTaskModel models.UpdateTaskModel) error {
  repo := service.CreateRepository()
	defer repo.Close()

	updateTask := mappers.UpdateTaskMapper.MapTo(updateTaskModel)
	updateTask.Status = entities.CommandNotReceived

	return repo.Add(updateTask)
}

func (service *updateService) SetUpdateStatus(updateTaskID uint, status entities.CommandStatus) {
	repo := service.CreateRepository()
	defer repo.Close()

	updateTask := &entities.UpdateTask{}
	updateTask.ID = updateTaskID

	repo.First(&updateTask, updateTask)

	updateTask.Status = status

	repo.Update(updateTask)
}

func (service *updateService) GetOutstandingTasks() []models.UpdateTaskModel {
	repo := service.CreateRepository()
	defer repo.Close()

	query := &entities.UpdateTask{Status: entities.CommandNotReceived}

	result := &[]entities.UpdateTask{}

	repo.GetWhere(&result, query)
	return mappers.UpdateTaskMapper.MultipleMapFrom(*result)
}

func (service *updateService) GetAllUpdates() []models.UpdateModel {
	repo := service.CreateRepository()
	defer repo.Close()

	result := &[]entities.Update{}
	repo.GetAll(result)

	return mappers.UpdateMapper.MultipleMapFrom(*result)
}

func (service *updateService) GetAllUpdateTasks() []models.UpdateTaskModel {
	repo := service.CreateRepository()
	defer repo.Close()

	result := []entities.UpdateTask{}
	repo.GetAll(result)

	return mappers.UpdateTaskMapper.MultipleMapFrom(result)
}

func (service *updateService) GetUpdateById(id uint) models.UpdateModel {
	repo := service.CreateRepository()
	defer repo.Close()

	result := &entities.Update{}
	query := &entities.Update{}
	query.ID = id

	repo.First(result, query)

	return mappers.UpdateMapper.MapFrom(*result)
}

func (service *updateService) GetUpdateTaskById(id uint) models.UpdateTaskModel {
	repo := service.CreateRepository()
	defer repo.Close()

	result := &entities.UpdateTask{}
	query := &entities.UpdateTask{}
	query.ID = id

	repo.First(result, query)

	return mappers.UpdateTaskMapper.MapFrom(*result)
}