package entities

type Update struct {
	GormEntity
	PayloadURL             string
	SuitableDeviceModel    DeviceModel
	SuitableDeviceModelID  int					
	SuitableVersion        float64
	CarMustStand           bool
	MinimalBattery         float64
}
