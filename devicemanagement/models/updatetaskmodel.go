package models

import (
	"time"

	"nova/devicemanagement/datastorage/entities"
)

type UpdateTaskModel struct {
	ModelBase
	DeviceID              int
	UpdateID              int
	SuitableDeviceModelID int
	PayloadURL            string
	MinimalBattery        float64
	CarMustStand          bool
	Status                entities.CommandStatus
	ReceivedAt            time.Time
	ExecutedAt            time.Time
}
