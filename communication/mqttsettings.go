package communication

import "nova/shared"

// shared constants between device and server
var CommandNotificationPathPattern = "/device/%v/commands/notify"
var CommandConfirmationPathPattern = "/device/%v/commands/statuschanged"
var UpdateTelemetryPathPattern = "/device/%v/telemetry/update"

type MQTTSettings struct {
	Server              string
	Port                int
	UseTLS              bool
	Username            string
	Password            string
	ClientID            string
	CertificateSettings *shared.CertificateSettings
}
