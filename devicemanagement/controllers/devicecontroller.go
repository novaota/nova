package controllers

import (
	"nova/devicemanagement/models"
	"nova/devicemanagement/services"
	"nova/rest"
)

type deviceController struct {
	services.DeviceService
}

func NewDeviceController(deviceService services.DeviceService) *deviceController {
	return &deviceController{deviceService}
}

func (controller *deviceController) GetRoutes() []rest.Route {
	r1 := &rest.Route{
		Method:                rest.MethodGet,
		Path:                  "/devices",
		RequiresAuthorization: true,
		Handler:               controller.GetAllDevices,
	}

	r2 := &rest.Route{
		Method:                rest.MethodPost,
		Path:                  "/devices",
		RequiresAuthorization: true,
		Handler:               controller.AddDevice,
	}

	return []rest.Route{*r1, *r2}
}

func (controller *deviceController) GetAllDevices(ctx *rest.RequestContext) (interface{}, rest.Error) {
	return controller.DeviceService.GetAll(), nil
}

func (controller *deviceController) AddDevice(ctx *rest.RequestContext) (interface{}, rest.Error) {
	device := &models.CarModel{}

	err := ctx.DecodeBody(&device)
	return device, err
}

func (controller *deviceController) RemoveDevices(ctx *rest.RequestContext) (interface{}, rest.Error) {
	return nil, nil
}
