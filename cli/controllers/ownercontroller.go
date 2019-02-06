package controllers

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/abiosoft/ishell.v2"
	"nova/devicemanagement/devicemgmtclient"
	"nova/devicemanagement/models"
)

func NewOwnerController(dmc devicemgmtclient.Client) *ownerController {
	return &ownerController{
		dmc: dmc,
	}
}

type ownerController struct {
	ControllerBase
	dmc devicemgmtclient.Client
}

func (c *ownerController) GetCommands() []CliCommand {
	c1 := CliCommand{
		Name: "addowner",
		Handler: c.AddOwner,
	}

	c2 := CliCommand{
		Name: "showowners",
		Handler: c.ShowAllOwners,
	}
	return []CliCommand{c1, c2}
}

func (c *ownerController) AddOwner(ctx *ishell.Context) {

	model := models.OwnerModel{
		Name: c.ReadString(ctx, "Name"),
		FirstName: c.ReadString(ctx, "FirstName"),
		Address: c.ReadString(ctx, "Address"),
		Zip: c.ReadString(ctx, "Zip"),
	}

	err := c.dmc.AddOwner(model)

	if err != nil {
		ctx.Printf("Error: %v \n", err.Error())
	} else {
		ctx.Println("Added")
	}
}

func (c *ownerController) ShowAllOwners(ctx *ishell.Context) {
	devices, err := c.dmc.GetAllOwners()

	if err != nil {
		ctx.Printf("Error: %v\n", err.Error())
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"FirstName",
		"Name",
		"Address",
		"Zip",
	})

	for _, owner := range devices {
		table.Append([]string {
			fmt.Sprintf("%v", owner.ID),
			owner.FirstName,
			owner.Name,
			owner.Address,
			owner.Zip,
		})
	}

	table.Render()
}