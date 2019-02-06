package main

import (
	"fmt"

	"nova/commands/update"
	"nova/communication"
)

func main() {

	config := CarConfig{
		DeviceId: 2,
	}

	fmt.Println("Command Engine")
	fmt.Println(" > Starting up")

	engine := NewCommandEngine(communication.DefaultMQTTSettings, config.DeviceId)

	//Register services
	updateHandler := update.NewUpdateHandler()

	handlers := []communication.CommandHandler{
		updateHandler,
	}

	fmt.Println(" > Register Handlers")

	for _, handler := range handlers {
		fmt.Printf(" > %v registered.\n", handler.GetIdentifier())
	}

	engine.Start()
}
