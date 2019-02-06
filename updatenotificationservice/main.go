package main

import (
	"nova/communication"
	"time"
)

func main() {
	n := NewNotificationService()

	mqs := &communication.MQTTSettings{Server: "localhost", Port: 12345}
	n.SetupMQTTSettings(mqs)
	n.SetupPostgresDatabase("host=localhost port=5432 user=postgres dbname=devicemanagement password=passwort sslmode=disable")
	n.SetupServices()
	n.SetSleepingInterval(10 * time.Second)

	n.Start()

}
