package controllers

import (
	"gopkg.in/abiosoft/ishell.v2"
)

type CliController interface {
	GetCommands() []CliCommand
}

type CliCommand struct {
	Name    string
	Help    string
	Handler func(ctx *ishell.Context)
}


type CommandLineInterface struct {
	Controllers []CliController
}

func (cli *CommandLineInterface) AddController(controllers ...CliController) {
  for _, controller := range controllers {
  	cli.Controllers = append(cli.Controllers, controller)
	}
}

func (cli *CommandLineInterface) Start() {
	shell := ishell.New()

	for _, controller := range cli.Controllers {
		for _, command := range controller.GetCommands() {
			shell.AddCmd(&ishell.Cmd{
				Name:      command.Name,
				Func:      command.Handler,
				Help:      command.Help,
			})
		}
	}

	shell.Run()
}