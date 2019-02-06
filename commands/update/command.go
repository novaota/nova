package update

import "nova/communication"

type Command struct {
	communication.Command
	DeviceID              int
	SuitableDeviceModelID int
	TaskID                uint
	PayloadURL            string
	MinimalBattery        float64
	CarMustStand          bool
}
