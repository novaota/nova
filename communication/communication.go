package communication

import (
	"fmt"
	"nova/devicemanagement/datastorage/entities"
)

func BuildNotificationTopic(device entities.Device) string {
	return fmt.Sprintf(CommandNotificationPathPattern, device.ID)
}

func BuildNotificationTopicById(deviceId int) string {
	return fmt.Sprintf(CommandNotificationPathPattern, deviceId)
}

func BuildConfirmationTopic(device entities.Device) string {
	return fmt.Sprintf(CommandConfirmationPathPattern, device.ID)
}

func BuildConfirmationTopicById(deviceId int) string {
	return fmt.Sprintf(CommandConfirmationPathPattern, deviceId)
}

func BuildGenericConfirmationTopic() string {
	return fmt.Sprintf(CommandConfirmationPathPattern, "+")
}

var DefaultMQTTSettings *MQTTSettings = &MQTTSettings{Server: "localhost", Port: 12345}
