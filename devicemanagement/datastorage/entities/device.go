package entities

import "time"

type Device struct {
	GormEntity
	SoftwareVersion float64
	DeviceModel     DeviceModel
	DeviceModelID   int
	Owner           Owner
	OwnerID         int
	ModelNr				  string
	LastConnection	time.Time
	CurrentBattery	float64
}
