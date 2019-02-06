package main

func main() {
	api := NewCarApi( 8080)
	api.SetName("CarInformationModel Simulation Api")
	api.StartServing()
}
