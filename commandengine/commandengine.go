package main

import (
	"log"
	"time"

	"nova/communication"
)

type CommandEngine struct {
	CommandListener *CommandListener
	CommandQueue    *CommandQueue
	handlers        []communication.CommandHandler
	BrokerService   communication.BrokerService
	isRunning			  bool
	deviceId		    int
}

func NewCommandEngine(settings *communication.MQTTSettings, deviceId int) *CommandEngine {
	return &CommandEngine{
		BrokerService: communication.NewMQTTBrokerService(settings),
		deviceId: deviceId,
	}
}

func (engine *CommandEngine) Start() {
	// Setting up command listener
	engine.CommandListener = NewCommandListener(engine.BrokerService, engine.deviceId)
	engine.CommandListener.receiver = engine

	// Setting up command queue
	engine.CommandQueue = NewCommandQueue()

	engine.isRunning = true

	engine.CommandListener.Start()
	engine.ExecuteOutstandingCommands()
}

func (engine *CommandEngine) ExecuteOutstandingCommands() {
	for engine.isRunning {

		if !engine.CommandQueue.HasCommands() {
			engine.waitForNewCommands()
		}

		//Zustandsanalyse
		//Überprüfen ob Zustand da ist

		for engine.CommandQueue.HasCommands() {
			currentCommand := engine.CommandQueue.Pop()
			engine.ExecuteCommand(currentCommand)
		}
	}
}

func (engine *CommandEngine) waitForNewCommands() {
	for !engine.CommandQueue.HasCommands() {
		time.Sleep(1*time.Second)
	}
}

func (engine *CommandEngine) ExecuteCommand(command *RawCommand) {
	log.Printf("CommandEngine: Executing command for Handler: with id %v\n", command.CommandHandler, command.ID)
	//TODO: Break here
	handler := engine.getCorrespondingHandler(command.CommandHandler)

	if handler == nil {
		log.Fatalf("Could not find handler for ")
		return
	}

	handler.SetByteParameters(command.Data)

	err := handler.Do()

	if !engine.isCommandHealthy(command) {
		engine.CommandQueue.RemoveCommand(command)
	}

	if err == nil {
		engine.CommandQueue.SetExecuted(command)
	}

	log.Fatalf("CommandEgine: Execution failed: %v\n", err.Error())

	err = handler.Undo()

	if err != nil {
		log.Fatalf("CommandEngine: Undo failed, no we have a serious problem")
	}

	engine.reevaluateCommandHealth(command)
}

func (engine *CommandEngine) Stop() {
	engine.isRunning = false
}

func (engine *CommandEngine) GetStatus() {
  //  How to implement status
}

func (engine *CommandEngine) RegisterHandler(handler communication.CommandHandler) {
	log.Printf("CommandEngine: Handler with Id: %v registered\n", handler.GetIdentifier())
	engine.handlers = append(engine.handlers, handler)
}

func (engine *CommandEngine) OnCommandReceived(command *RawCommand) bool{
	log.Printf("CommandEngine: Received Command for Handler: %v with id %v\n", command.CommandHandler, command.ID)
	engine.CommandQueue.Push(command)
	return true
}

func (engine *CommandEngine) getCorrespondingHandler(handlerName string) communication.CommandHandler{
	for _, handler := range engine.handlers {
		if handler.GetIdentifier() == handlerName {
			return handler
		}
	}
	return nil
}

func (engine *CommandEngine) isCommandHealthy (command *RawCommand) bool {
  tries := engine.CommandQueue.getCommandoTries(command)
  return tries <= command.MaxTries
}

func (engine *CommandEngine) reevaluateCommandHealth(command *RawCommand) {
	engine.CommandQueue.incrementCommandoTries(command)
}