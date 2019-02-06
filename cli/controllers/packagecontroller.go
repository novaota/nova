package controllers

import (
	"fmt"
	"os"
	"gopkg.in/abiosoft/ishell.v2"
	"nova/shared"
	"nova/updatepackage"
)

func NewPackController(packer *updatepackage.Packer, unpacker *updatepackage.Unpacker) *packController{
	return &packController{
		packer: packer,
		unpacker: unpacker,
	}
}

type packController struct {
	packer *updatepackage.Packer
	unpacker *updatepackage.Unpacker
	ControllerBase
}

func (c *packController) GetCommands() []CliCommand {
	c1 := CliCommand{
		Name:    "pack",
		Handler: c.Pack,
	}

	c2 := CliCommand{
		Name:    "unpack",
		Handler: c.Unpack,
	}

	return []CliCommand{c1, c2}
}

func (c *packController) Pack(ctx *ishell.Context) {
	ctx.ShowPrompt(false)
	defer ctx.ShowPrompt(true)

	name := c.ReadString(ctx, "Name")
  method := c.ReadString(ctx, "Method")

	files := c.gatherFiles(ctx)
	outputDir := c.chooseDir(ctx)

	up := &updatepackage.UpdatePackage{
		Name: name,
		Files: files,
		UpdateMethod: method,
	}

	_, err := c.packer.Pack(up, outputDir)

	if err != nil {
		ctx.Println(fmt.Sprintf("Could not create package %v.", err.Error()))
	}
}

func (c *packController) Unpack(ctx *ishell.Context) {
	ctx.ShowPrompt(false)
	defer ctx.ShowPrompt(true)

	ctx.Print("Deploymentfile:")
	deploymentFile := ctx.ReadLine()

	deployment := &updatepackage.UpdatePackageDeployment{}

	err := shared.DefaultSerializer.DeserializeFromFile(deploymentFile, deployment)

	if err != nil {
		ctx.Println("Invalid deployment file")
	}

	outputDir := c.chooseDir(ctx)

	_, err = c.unpacker.Unpack(*deployment, outputDir)

	if err != nil {
		ctx.Printf("Error: %v\n" ,err.Error())
	}
}

func (c *packController) chooseDir(ctx *ishell.Context) string  {
	ctx.Print("Output Directory:")
	dir := ctx.ReadLine()

	if !shared.FileOrFolderExists(dir) {
		os.Mkdir(dir, os.ModeDir)
	}

	return dir
}

func (c *packController) gatherFiles(ctx *ishell.Context) []string {
	var files []string
	for {
		ctx.Print("Add File: ")
		file := ctx.ReadLine()

		if len(file) == 0 {
			break
		}

		files = append(files, file)
	}
	return files
}
