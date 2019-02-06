package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"nova/communication"
	"nova/shared"
)

type CommandQueueItem struct {
	Command    *RawCommand
	Executed   bool
	ReceivedAt time.Time
	Started    bool
	Ended      bool
	Tries      uint
}

type CommandQueue struct {
	queue []CommandQueueItem
}

func NewCommandQueue() *CommandQueue {
	return &CommandQueue{}
}

func (cq *CommandQueue) Push(command *RawCommand) {
	item := &CommandQueueItem{
		Command:  command,
		Executed: false,
	}
	cq.queue = append(cq.queue, *item)
	cq.Save()
}

func (cq *CommandQueue) Pop() *RawCommand {
	q := cq.queue

	if len(q) < 1 {
		return nil
	}

	return q[len(q)-1].Command
}

func (cq *CommandQueue) SetExecuted(command *RawCommand) {
	item := cq.getItemByRawCommand(command)
	item.Executed = true
	cq.Save()
}

func (cq *CommandQueue) Save() {
	log.Println("CommandPersistende: Saving current state..")

	path := cq.getDbStorePath()

	data := cq.queue

	err := shared.DefaultSerializer.SerializeToFile(data, path)

	if err != nil {
		log.Fatalln("CommandPersistence: Could not save commands")
	}
}

func (cq *CommandQueue) Restore() {
	log.Println("CommandPersistende: Restoring database")
	path := cq.getDbStorePath()

	queue := &[]CommandQueueItem{}
	err := shared.DefaultSerializer.DeserializeFromFile(path, queue)

	if err != nil {
		log.Fatalf("CommandPersistence: Could not restore from file %v. \n", path)
	}

	cq.queue = *queue
}

func (cq *CommandQueue) HasCommands() bool {
	return len(cq.queue) > 0
}

func (cq *CommandQueue) incrementCommandoTries(command *RawCommand) {
  item := cq.getItemByRawCommand(command)

	if item == nil {
		return
	}

  item.Tries++
}

func (cq *CommandQueue) getCommandoTries(command *RawCommand) uint {
	item := cq.getItemByRawCommand(command)

	return item.Tries
}

func (cq *CommandQueue) getDbStorePath() string {
	dbName := "commands.json"
	executablePath := os.Args[0]
	dir := filepath.Dir(executablePath)
	return filepath.Join(dir, dbName)
}

func (cq *CommandQueue) getItemByCommand(command *communication.Command) *CommandQueueItem {
	for _, item := range cq.queue {
		if item.Command.ID == command.ID {
			return &item
		}
	}
	return nil
}

func (cq *CommandQueue) getItemByRawCommand(command *RawCommand) *CommandQueueItem {
	for _, item := range cq.queue {
		if item.Command.ID == command.ID {
			return &item
		}
	}
	return nil
}

func (cq *CommandQueue) RemoveCommand(command *RawCommand) {
	for i, item := range cq.queue {
		if item.Command.ID == command.ID {
			cq.queue = append(cq.queue[:i], cq.queue[i+1:]...)
			return
		}
	}
}
