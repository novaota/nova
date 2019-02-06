package main

import (
	"encoding/json"
	"fmt"

	"nova/communication"
	"nova/devicemanagement/datastorage/entities"
)

type CommandReceiver interface {
	OnCommandReceived(command *RawCommand) bool
}

type CommandListener struct {
	receiver      CommandReceiver
	BrokerService communication.BrokerService
	DeviceId      int
}

type RawCommand struct {
	Data           []byte
	CommandHandler string
	MaxTries       uint
	ID             uint
}

func NewCommandListener(brokerService communication.BrokerService, deviceId int) *CommandListener {
	return &CommandListener{
		BrokerService: brokerService,
		DeviceId: deviceId,
	}
}

func (listener *CommandListener) Start() {
	listener.BrokerService.Connect()
	topic := communication.BuildNotificationTopicById(listener.DeviceId)
	listener.BrokerService.Subscribe(topic, listener.newCommandArrived)
}

func (listener *CommandListener) newCommandArrived(topic string, data []byte) {
	command := &communication.Command{}
	err := json.Unmarshal(data, command)

	strData := string(data)

	fmt.Println(strData)

	if err != nil {
		return
	}

	handler := command.CommandHandler
	id := command.ID
  maxTries := command.MaxTries

	if listener.receiver == nil {
		return
	}

	rawCommand := &RawCommand{
		Data: data,
		CommandHandler: handler,
		MaxTries: maxTries,
		ID: id,
	}

	isHandled := listener.receiver.OnCommandReceived(rawCommand)


	if isHandled {
		listener.confirmReceivedCommand(command)
	}
}

func (listener *CommandListener) confirmReceivedCommand(command *communication.Command) {
	listener.sendConfirmationMessage(command, entities.CommandReceived)
}

func (listener *CommandListener) confirmExecutedCommand(command *communication.Command) {
	listener.sendConfirmationMessage(command, entities.CommandExecuted)
}

func (listener *CommandListener) sendConfirmationMessage(command *communication.Command, status entities.CommandStatus) {
	topic := communication.BuildConfirmationTopicById(listener.DeviceId)

	confirmation := communication.CommandConfirmation{
		ID: command.ID,
		Status:    status,
	}

	listener.BrokerService.Publish(confirmation, topic)
}