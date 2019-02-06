package models

type UpdateModel struct {
	ModelBase
	PayloadURL            string
	SuitableDeviceModelID int
	SuitableVersion       float64
	CarMustStand          bool
	MinimalBattery        float64
}
