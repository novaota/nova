package main

import (
	"log"
	"math"
	"time"

	"nova/car/models"
)

var version float64 = 1.1

type CarSimulation struct {
	models.CarInformationModel
	IsSimulating bool
}

func NewCarSimulation() *CarSimulation {
	simulation := &CarSimulation{}
	simulation.Start()
	return simulation
}

// properties
func (sim *CarSimulation) Start() {

	log.Println("CarSimulation: Start")

	sim.IsSimulating = true
	sim.Version = version

	sim.NetworkReception = 0.5
	sim.LoadBattery()

	go sim.SimulateBatteryBehaviour()
	go sim.SimulateChangingNetworkReception()
}

func (sim *CarSimulation) StartDriving() {
	sim.IsDriving = true
}

func (sim *CarSimulation) StopDriving() {
	sim.IsDriving = false
}

func (sim *CarSimulation) LoadBattery() {
	sim.BatteryLevel = 1
}

func (sim *CarSimulation) SimulateBatteryBehaviour() {
	var i float64
	sim.BatteryLevel = 1
	for sim.IsSimulating {
		for i = sim.BatteryLevel; i >= 0; i -= 0.01 {
			sim.BatteryLevel = i
			time.Sleep(500 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (sim *CarSimulation) SimulateChangingNetworkReception() {
	var multiplier float64 = 1.0
	var step float64 = 0.05

	networkReception := 0.5

	for sim.IsSimulating {
		if networkReception >= 1.0 {
			multiplier = -1.0
		} else if networkReception <= -0.2 {
			multiplier = 1.0
		}

		networkReception += (multiplier * step)

		sim.NetworkReception = math.Max(0.0, networkReception)
		time.Sleep(500 * time.Millisecond)
	}
}