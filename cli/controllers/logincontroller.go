package controllers

import (
	"gopkg.in/abiosoft/ishell.v2"
	"nova/devicemanagement/devicemgmtclient"
)

func NewLoginController(dmc devicemgmtclient.Client) *loginController{
	return &loginController{
		dmc: dmc,
	}
}

type loginController struct {
	dmc devicemgmtclient.Client
}

func (c *loginController) GetCommands() []CliCommand {
	c1 := CliCommand{
		Name:"login",
		Handler: c.Login,
	}

	return []CliCommand{c1,	}
}

func (c *loginController) Login(ctx *ishell.Context) {
	ctx.ShowPrompt(false)
	defer ctx.ShowPrompt(true)

	ctx.Print("Username: ")
	username := ctx.ReadLine()

	// get password.
	ctx.Print("Password: ")
	password := ctx.ReadPassword()

	err := c.dmc.Login(username, password)

	if err == nil {
		ctx.Print("Logged In Successfully")//
	} else {
		ctx.Print("Login failed")
	}
}

