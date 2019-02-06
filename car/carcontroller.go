package main

import (
	"nova/rest"
)

type CarController struct {
	Car *CarSimulation
}

func NewSimulatedCarController() *CarController {
	return &CarController{
		Car: NewCarSimulation(),
	}
}

func (controller *CarController) GetRoutes() []rest.Route {

	r1 := rest.Route{Method: rest.MethodGet,
		Path:                  "/car/driving/start",
		RequiresAuthorization: false,
		Handler:               controller.StartDriving,
	}

	r2 := rest.Route{Method: rest.MethodGet,
		Path:                  "/car/driving/stop",
		RequiresAuthorization: false,
		Handler:               controller.StopDriving,
	}

	r3 := rest.Route{Method: rest.MethodGet,
		Path:                  "/car/status",
		RequiresAuthorization: false,
		Handler:               controller.GetStatus,
	}

	r4 := rest.Route{Method: rest.MethodGet,
		Path:                  "/car/recharge",
		RequiresAuthorization: false,
		Handler:               controller.LoadBattery,
	}

	return []rest.Route{r1, r2, r3, r4}
}

func (controller *CarController) StartDriving(ctx *rest.RequestContext) (interface{}, rest.Error) {
	controller.Car.StartDriving()
	return nil, nil
}

func (controller *CarController) StopDriving(ctx *rest.RequestContext) (interface{}, rest.Error) {
	controller.Car.StopDriving()
	return nil, nil
}

func (controller *CarController) GetStatus(ctx *rest.RequestContext) (interface{}, rest.Error) {
	return controller.Car, nil
}

func (controller *CarController) LoadBattery(ctx *rest.RequestContext) (interface{}, rest.Error) {
	controller.Car.LoadBattery()
	return nil, nil
}
