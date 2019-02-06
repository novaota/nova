package communication

type UpdateCommand struct {
	Command
	ID                    uint
	DeviceID              int
	PayloadURL            string
	SourceVersion         float64
	TargetVersion	        float64
	UpdateRequirements
}

type UpdateRequirements struct {
	CarMustStand            bool
	MinimalBattery          float64
	MinimalNetworkReception float64
	SuitableDeviceModelID   int
}