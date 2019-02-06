package communication

import "nova/devicemanagement/datastorage/entities"


type Command struct {
	ID                  uint
	MaxTries			      uint
	CommandVersion      float64
	CommandHandler      string
	Commands            []Command
}

type CommandHandler interface {
	GetIdentifier() string
	SetParameters(command interface{})
	SetByteParameters(data []byte)
	Do() error
	Undo() error
	IsExecutable() bool
}

type CommandConfirmation struct {
	ID        uint
	DeviceID  uint
	Status    entities.CommandStatus
}