package models

type CarInformationModel struct {
	DeviceId         int               `json:"device_id"`
	IsDriving        bool              `json:"is_driving"`
	BatteryLevel     float64           `json:"battery_level"`
	NetworkReception float64           `json:"network_reception"`
	Version          float64           `json:"version"`
}
