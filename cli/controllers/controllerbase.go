package controllers

import (
	"strconv"
	"gopkg.in/abiosoft/ishell.v2"
)

type ControllerBase struct { }

func (c *ControllerBase) ReadString(ctx *ishell.Context, name string) string {
	ctx.ShowPrompt(false)
	defer ctx.ShowPrompt(true)

	ctx.Print(name + ":")
	return ctx.ReadLine()
}

func (c *ControllerBase) ReadInt(ctx *ishell.Context, name string) int {
	for {
		str := c.ReadString(ctx, name)
		value, err :=  strconv.Atoi(str)

		if err == nil {
			return value
		}

		ctx.Println("Invalid")
	}
}

func (c *ControllerBase) ReadFloat64(ctx *ishell.Context, name string) float64 {
	for {
		str := c.ReadString(ctx, name)
		value, err :=  strconv.ParseFloat(str, 64)

		if err == nil {
			return value
		}

		ctx.Println("Invalid")
	}
}

func (c *ControllerBase) ReadFloat32(ctx *ishell.Context, name string) float32 {
	f64 := c.ReadFloat64(ctx, name)
	return float32(f64)
}

func (c *ControllerBase) ReadBool(ctx *ishell.Context, name string) bool {
	choice := ctx.MultiChoice([]string{"Yes", "No"}, name)
	return choice == 0
}