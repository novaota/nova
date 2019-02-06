package main

import (
	"encoding/json"
	"log"
	"nova/communication"
	"nova/devicemanagement/datastorage"
	"nova/devicemanagement/services"
	"nova/devicemanagement/services/implementation"
	"time"
)

type ConfirmationService struct {
	RepositoryFactory datastorage.RepositoryFactory
	UpdateService     services.UpdateService
	BrokerService     communication.BrokerService
	MQTTSettings      *communication.MQTTSettings
}

func NewConfirmationService() *ConfirmationService {
	return &ConfirmationService{}
}

// Setup Methods

func (service *ConfirmationService) SetupMQTTSettings(settings *communication.MQTTSettings) {
	service.MQTTSettings = settings
}

func (service *ConfirmationService) SetupServices() {
	service.BrokerService = communication.NewMQTTBrokerService(service.MQTTSettings)
	service.UpdateService = implementation.NewUpdateService(service.RepositoryFactory)
}

// public setup methods
func (service *ConfirmationService) SetupPostgresDatabase(connectionString string) {
	service.RepositoryFactory = datastorage.NewPostgresRepository(connectionString)
}

func (service *ConfirmationService) Start() {
	log.Println("Starting UpdateNotificationService")

	err := service.BrokerService.Connect()

	if err != nil {
		panic(err)
	}

	topic := communication.BuildGenericConfirmationTopic()
	service.BrokerService.Subscribe(topic, service.newConfirmationArrived)

	go forever()
}

func (service *ConfirmationService) newConfirmationArrived(topic string, data []byte) {
	confirmation := &communication.CommandConfirmation{}
	err := json.Unmarshal(data, &confirmation)

	if err != nil {
		log.Fatalf("Could not parse incoming confirmation content: %v\n", string(data))
	}

	log.Printf("Updates Status of Device %v.", confirmation.DeviceID)
	service.UpdateService.SetUpdateStatus(confirmation.ID, confirmation.Status)
}

func forever() {
	for {
		time.Sleep(1 * time.Second)
	}
}
