package entities

import "time"

type UpdateTask struct {
	GormEntity
	Device     Device
	DeviceID   int
	Update     Update
	UpdateID   int
	Status     CommandStatus
	ReceivedAt time.Time
	ExecutedAt time.Time
}
