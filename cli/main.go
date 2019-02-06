package main

import (
	"fmt"

	"nova/cli/controllers"
	"nova/devicemanagement/devicemgmtclient"
	"nova/shared"
	"nova/updatepackage"
)

func printHeader() {
	fmt.Println("  _   _    ____   __      __                  _____   _        _____ ")
	fmt.Println(" | \\ | |  / __ \\  \\ \\    / /     /\\          / ____| | |      |_   _|")
	fmt.Println(" |  \\| | | |  | |  \\ \\  / /     /  \\        | |      | |        | |  ")
	fmt.Println(" | . ` | | |  | |   \\ \\/ /     / /\\ \\       | |      | |        | |  ")
	fmt.Println(" | |\\  | | |__| |    \\  /     / ____ \\      | |____  | |____   _| |_ ")
	fmt.Println(" |_| \\_|  \\____/      \\/     /_/    \\_\\      \\_____| |______| |_____|")
	fmt.Println("                                                                     ")
}

func main() {

	printHeader()

	dmc := *devicemgmtclient.New(shared.CertificateSettings{}, "http://localhost:8080/")

	loginController := controllers.NewLoginController(dmc)
	carModelController := controllers.NewCarModelController(dmc)
	deviceController := controllers.NewDeviceController(dmc)
	ownerController := controllers.NewOwnerController(dmc)
	updateController := controllers.NewUpdateController(dmc)
	updateTaskController := controllers.NewUpdateTaskController(dmc)

	packer := updatepackage.NewPacker(shared.DefaultCertificateSettings)
	unpacker := updatepackage.NewUnpacker(shared.DefaultCertificateSettings)

	packageController := controllers.NewPackController(packer, unpacker)

	cli := &controllers.CommandLineInterface{}
	cli.AddController(loginController,
		                carModelController,
		                deviceController,
		                ownerController,
		                updateController,
		                updateTaskController,
		                packageController)

	cli.Start()
}



