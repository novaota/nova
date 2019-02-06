package models

import "time"

type CarModel struct {
	ModelBase
	SoftwareVersion float64
	DeviceModelID   int
	OwnerID         int
	ModelNr         string
	LastConnection  time.Time
	CurrentBattery  float64
}
