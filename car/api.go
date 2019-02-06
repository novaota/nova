package main

import (
	"nova/rest"
)

type CarApi struct {
	*rest.Api
}

func NewCarApi(port int) *CarApi {
	api := &CarApi{
		Api: rest.NewRestApi(port),
	}

	controller := NewSimulatedCarController()
	api.AddController(controller)

	return api
}
