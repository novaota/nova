package entities

type CommandStatus int

const (
	CommandNotReceived CommandStatus = 0
	CommandReceived    CommandStatus = 1
	CommandExecuted    CommandStatus = 2
)
