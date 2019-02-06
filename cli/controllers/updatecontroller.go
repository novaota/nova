package controllers

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/abiosoft/ishell.v2"
	"nova/devicemanagement/devicemgmtclient"
	"nova/devicemanagement/models"
)

func NewUpdateController(dmc devicemgmtclient.Client) *updateController {
	return &updateController{
		dmc: dmc,
	}
}

type updateController struct {
	ControllerBase
	dmc devicemgmtclient.Client
}

func (c *updateController) GetCommands() []CliCommand {
	c1 := CliCommand{
		Name: "addupdate",
		Handler: c.AddUpdate,
	}

	c2 := CliCommand{
		Name: "showupdates",
		Handler: c.ShowAllUpdates,
	}
	return []CliCommand{c1, c2}
}

func (c *updateController) AddUpdate(ctx *ishell.Context) {
	model := models.UpdateModel{
		MinimalBattery: c.ReadFloat64(ctx, "MinimalBattery"),
		CarMustStand: c.ReadBool(ctx, "CarMustStand"),
		PayloadURL: c.ReadString(ctx, "PayloadURL"),
		SuitableDeviceModelID: c.ReadInt(ctx, "SuitableDeviceModelId"),
		SuitableVersion: c.ReadFloat64(ctx, "SuitableVersion"),
	}

	err := c.dmc.AddUpdate(model)

	if err != nil {
		ctx.Printf("Error: %v \n", err.Error())
	} else {
		ctx.Println("Added")
	}
}

func (c *updateController) ShowAllUpdates(ctx *ishell.Context) {
	updates, err := c.dmc.GetAllUpdates()

	if err != nil {
		ctx.Printf("Error: %v\n", err.Error())
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"SuitableVersion",
		"SuitableDeviceModelId",
		"PayloadUrl",
		"CarMustStand",
		"MinimalBattery",
	})

	for _, update := range updates {
		table.Append([]string {
			fmt.Sprintf("%v", update.ID),
			fmt.Sprintf("%f", update.SuitableVersion),
			fmt.Sprintf("%v", update.SuitableDeviceModelID),
			update.PayloadURL,
			fmt.Sprintf("%b", update.CarMustStand),
			fmt.Sprintf("%f", update.MinimalBattery),
		})
	}

	table.Render()
}