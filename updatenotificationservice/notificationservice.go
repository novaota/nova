package main

import (
	"log"
	"time"

	"nova/commands/update"
	"nova/communication"
	"nova/devicemanagement/datastorage"
	"nova/devicemanagement/models"
	"nova/devicemanagement/services"
	"nova/devicemanagement/services/implementation"
)

type NotificationService struct {
	MQTTSettings *communication.MQTTSettings

	// DataStorage
	RepositoryFactory datastorage.RepositoryFactory

	// Controller
	BrokerService communication.BrokerService
	UpdateService services.UpdateService

	//running flag
	isRunning     bool
	sleepTimeSpan time.Duration
}

// initialization methods
func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (service *NotificationService) Start() {
	log.Println("Starting Notification Service")

	if service.isRunning {
		log.Fatalln(" > Could not start NotificationService, because its already running")
	}

	service.isRunning = true

	log.Println(" > Starting NotificationService")
	log.Printf(" > Notify every %v", service.sleepTimeSpan)

	service.run()
}

func (service *NotificationService) SetSleepingInterval(duration time.Duration) {
	service.sleepTimeSpan = duration
}

func (service *NotificationService) Stop() {
	service.isRunning = false
}

func (service *NotificationService) run() {
	service.BrokerService.Connect()

	for service.isRunning {
		service.NotifyDevices()
		time.Sleep(service.sleepTimeSpan)
	}
}

func (service *NotificationService) SetupMQTTSettings(settings *communication.MQTTSettings) {
	service.MQTTSettings = settings
}

func (service *NotificationService) SetupServices() {
	service.BrokerService = communication.NewMQTTBrokerService(service.MQTTSettings)
	service.UpdateService = implementation.NewUpdateService(service.RepositoryFactory)
}

// public setup methods
func (service *NotificationService) SetupPostgresDatabase(connectionString string) {
	service.RepositoryFactory = datastorage.NewPostgresRepository(connectionString)
}

func (service *NotificationService) NotifyDevices() {
	outstandingTasks := service.UpdateService.GetOutstandingTasks()
	for _, task := range outstandingTasks {
		service.notifyDevice(task)
	}
}

func (service *NotificationService) notifyDevice(task models.UpdateTaskModel) {
	log.Printf("Sending update notification to device with id %v\n", task.DeviceID)
	topic := communication.BuildNotificationTopicById	(task.DeviceID)
	updateNotification := service.buildNotification(task)

	err := service.BrokerService.Publish(updateNotification, topic)
	if err != nil {
		log.Println(" > Could not send")
		return
	}

	log.Println(" > Sent")
}

func (service *NotificationService) buildNotification(task models.UpdateTaskModel) *update.Command {
	result := &update.Command{
		DeviceID:              task.DeviceID,
		SuitableDeviceModelID: task.SuitableDeviceModelID,
		TaskID:                task.ID,
		PayloadURL:            task.PayloadURL,
		MinimalBattery:        task.MinimalBattery,
		CarMustStand:          task.CarMustStand,
		// TODO: Nicht schön aber fürs erste reichts
		Command: communication.Command{
			CommandHandler: update.Identifier,
			MaxTries: 3,
			CommandVersion: 1.0,
			ID: task.ID,
		},
	}
	return result
}
