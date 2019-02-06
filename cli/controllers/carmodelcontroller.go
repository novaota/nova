package controllers

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/abiosoft/ishell.v2"
	"nova/devicemanagement/devicemgmtclient"
	"nova/devicemanagement/models"
)

func NewCarModelController(dmc devicemgmtclient.Client) *carModelController {
	return &carModelController{
		dmc: dmc,
	}
}

type carModelController struct {
	ControllerBase
	dmc devicemgmtclient.Client
}

func (c *carModelController) GetCommands() []CliCommand {
	c1 := CliCommand{
		Name: "addcarmodel",
		Handler: c.AddCarModel,
	}
	c2 := CliCommand{
		Name: "showcarmodels",
		Handler: c.ShowAllCarModels,
	}
	return []CliCommand{c1, c2}
}

func (c *carModelController) AddCarModel(ctx *ishell.Context) {

	model := models.CarModelModel{
		Name: c.ReadString(ctx, "Name"),
	}

	err := c.dmc.AddCarModel(model)

	if err != nil {
		ctx.Printf("Error: %v \n", err.Error())
	} else {
		ctx.Println("Added")
	}
}

func (c *carModelController) ShowAllCarModels(ctx *ishell.Context) {
	carModels, err := c.dmc.GetAllCarModels()

	if err != nil {
		ctx.Printf("Error: %v\n", err.Error())
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"Name",
	})

	for _, carModel := range carModels {
		table.Append([]string {
			fmt.Sprintf("%v", carModel.ID),
			carModel.Name,
		})
	}

	table.Render()
}