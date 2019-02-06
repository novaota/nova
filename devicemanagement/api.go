package main

import (
	"log"

	"nova/devicemanagement/controllers"
	"nova/devicemanagement/datastorage"
	"nova/devicemanagement/datastorage/entities"
	"nova/devicemanagement/services"
	"nova/devicemanagement/services/implementation"
	"nova/rest"
)

type DeviceManagement struct {
	// Api
	*rest.Api

	// DataStorage
	RepositoryFactory datastorage.RepositoryFactory

	// Services
	AuthenticationService    services.AuthenticationService
	StateService             services.StateService
	DeviceService            services.DeviceService
	UpdateService            services.UpdateService
	DeviceModelService       services.DeviceModelService
	OwnerService						 services.OwnerService
}

func NewDeviceManagementApi(port int) *DeviceManagement {
	return &DeviceManagement{
		Api: rest.NewRestApi(port),
	}
}

func (dm *DeviceManagement) StartServing() {
	dm.setupServices()
	dm.setupControllers()

	rest.StateService = dm.StateService

	dm.initializeDatabaseIfNeeded()
	dm.Api.StartServing()
}

func (dm *DeviceManagement) setupServices() {
	log.Println(" > Initializing Services")
	dm.AuthenticationService = implementation.NewDatabaseAuthenticationService(dm.RepositoryFactory)
	dm.StateService =          implementation.NewSessionStateService()
	dm.DeviceService =         implementation.NewDeviceService(dm.RepositoryFactory)
	dm.UpdateService =         implementation.NewUpdateService(dm.RepositoryFactory)
	dm.DeviceModelService =    implementation.NewDeviceModelService(dm.RepositoryFactory)
	dm.OwnerService =          implementation.NewOwnerService(dm.RepositoryFactory)
}

func (dm *DeviceManagement) setupControllers() {
	log.Println(" > Setting Up Controllers")

	loginController :=       controllers.NewLoginController(dm.AuthenticationService, dm.StateService)
	deviceController :=      controllers.NewDeviceController(dm.DeviceService)
	updateController :=      controllers.NewUpdateController(dm.UpdateService)
	updateTaskController :=  controllers.NewUpdateTaskController(dm.UpdateService)
	deviceModelController := controllers.NewDeviceModelController(dm.DeviceModelService)
	ownerController :=       controllers.NewOwnerController(dm.OwnerService)
	dm.AddController(loginController,
		               deviceController,
		               updateController,
		               updateTaskController,
		               deviceModelController,
		               ownerController)
}

func (dm *DeviceManagement) initializeDatabaseIfNeeded() {
	repo := dm.RepositoryFactory.CreateRepository()
	defer repo.Close()

	if !repo.IsSchemaCreated() {
		log.Println("Creating Database")
		repo.CreateSchema()
		log.Println(" > Inserting initial data")
		dm.InsertDummyData()
	}
}

// public setup methods
func (dm *DeviceManagement) SetupPostgresDatabase(connectionString string) {
	dm.RepositoryFactory = datastorage.NewPostgresRepository(connectionString)
}

func (dm *DeviceManagement) InsertDummyData() {
	repo := dm.RepositoryFactory.CreateRepository()
	defer repo.Close()

	owner := entities.Owner{
		Name:      "Almesberger",
		FirstName: "Felix",
		Address:   "Lorenzerstraße 18a",
		Zip:       "93138",
	}

	deviceModel := entities.DeviceModel{
		Name: "BMW 8000",
	}

	device := entities.Device{
		CurrentBattery:  1,
		DeviceModel:     deviceModel,
		ModelNr:         "123345",
		SoftwareVersion: 0.1,
		Owner:           owner,
	}

	update := entities.Update{
		SuitableVersion:     1.0,
		SuitableDeviceModel: deviceModel,
		PayloadURL:          "https://deploymentserver/blablablub.json",
	}

	updateTask := entities.UpdateTask{
		Device: device,
		Status: entities.CommandNotReceived,
		Update: update,
	}

	repo.Add(&owner)
	repo.Add(&deviceModel)
	repo.Add(&device)
	repo.Add(&update)
	repo.Add(&updateTask)

	dm.AuthenticationService.RegisterUser("felixalmesberger", "felixalmesberger")
}
