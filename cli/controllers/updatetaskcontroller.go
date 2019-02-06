package controllers

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/abiosoft/ishell.v2"
	"nova/devicemanagement/devicemgmtclient"
	"nova/devicemanagement/models"
)

func NewUpdateTaskController(dmc devicemgmtclient.Client) *updateTaskController {
	return &updateTaskController{
		dmc: dmc,
	}
}

type updateTaskController struct {
	ControllerBase
	dmc devicemgmtclient.Client
}

func (c *updateTaskController) GetCommands() []CliCommand {
	c1 := CliCommand{
		Name: "addupdatetask",
		Handler: c.AddUpdateTask,
	}

	c2 := CliCommand{
		Name: "showupdatetasks",
		Handler: c.ShowAllUpdatesTasks,
	}
	return []CliCommand{c1, c2}
}

func (c *updateTaskController) AddUpdateTask(ctx *ishell.Context) {
	model := models.UpdateTaskModel{
    DeviceID: c.ReadInt(ctx, "DeviceID"),
    UpdateID: c.ReadInt(ctx, "UpdateID"),
	}

	err := c.dmc.AddUpdateTask(model)

	if err != nil {
		ctx.Printf("Error: %v \n", err.Error())
	} else {
		ctx.Println("Added")
	}
}

func (c *updateTaskController) ShowAllUpdatesTasks(ctx *ishell.Context) {
	updateTasks, err := c.dmc.GetAllUpdateTasks()

	if err != nil {
		ctx.Printf("Error: %v\n", err.Error())
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"DeviceID",
		"UpdateID",
		"Status",
	})

	for _, updateTask := range updateTasks {
		table.Append([]string {
			fmt.Sprintf("%v", updateTask.ID),
			fmt.Sprintf("%v", updateTask.DeviceID),
			fmt.Sprintf("%v", updateTask.UpdateID),
			fmt.Sprintf("%v", updateTask.Status),
		})
	}

	table.Render()
}