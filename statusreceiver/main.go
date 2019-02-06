package main

import (
	"nova/communication"
)

func main() {
	s := NewConfirmationService()
	s.SetupMQTTSettings(communication.DefaultMQTTSettings)
	s.SetupPostgresDatabase("host=localhost port=5432 user=postgres dbname=devicemanagement password=passwort sslmode=disable")
	s.SetupServices()
	s.Start()
}
