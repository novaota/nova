package controllers

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/abiosoft/ishell.v2"
	"nova/devicemanagement/devicemgmtclient"
	"nova/devicemanagement/models"
)

func NewDeviceController(dmc devicemgmtclient.Client) *deviceController {
	return &deviceController{
		dmc: dmc,
	}
}

type deviceController struct {
	ControllerBase
	dmc devicemgmtclient.Client
}

func (c *deviceController) GetCommands() []CliCommand {
	c1 := CliCommand{
		Name: "adddevice",
		Handler: c.AddDevice,
	}

	c2 := CliCommand{
		Name: "showdevices",
		Handler: c.ShowAllDevices,
	}
	return []CliCommand{c1, c2}
}

func (c *deviceController) AddDevice(ctx *ishell.Context) {
	model := models.CarModel{
		OwnerID: c.ReadInt(ctx, "OwnerID"),
		DeviceModelID: c.ReadInt(ctx, "DeviceModelID"),
		ModelNr: c.ReadString(ctx, "ModelNr"),
	}

	err := c.dmc.AddDevice(model)

	if err != nil {
		ctx.Printf("Error: %v \n", err.Error())
	} else {
		ctx.Println("Added")
	}
}

func (c *deviceController) ShowAllDevices(ctx *ishell.Context) {
	devices, err := c.dmc.GetAllDevices()

	if err != nil {
		ctx.Printf("Error: %v\n", err.Error())
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"OwnerID",
		"DeviceModelID",
		"ModelNr",
		"SoftwareVersion",
		"CurrentBattery",
		"LastConnection",
	})

	for _, device := range devices {
		table.Append([]string {
			fmt.Sprintf("%v", device.ID),
			fmt.Sprintf("%v", device.OwnerID),
			fmt.Sprintf("%v", device.DeviceModelID),
			device.ModelNr,
			fmt.Sprintf("%f", device.SoftwareVersion),
			fmt.Sprintf("%f", device.CurrentBattery),
			device.LastConnection.Format("20060102150405"),
		})
	}

	table.Render()
}