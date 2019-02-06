package controllers

import (
	"net/http"

	"nova/devicemanagement/models"
	"nova/devicemanagement/services"
	"nova/rest"
)

func NewDeviceModelController(deviceModelService services.DeviceModelService) *deviceModelController {
	return &deviceModelController{
		deviceModelService: deviceModelService,
	}
}

type deviceModelController struct {
	deviceModelService services.DeviceModelService
}

func (controller *deviceModelController) GetRoutes() []rest.Route {
	r1 := rest.Route{
		Method:                rest.MethodGet,
		Path:                  "/devices/models",
		RequiresAuthorization: false,
		Handler:               controller.getAll,
	}

	r2 := rest.Route{
		Method:                rest.MethodPost,
		Path:                  "/devices/models",
		RequiresAuthorization: false,
		Handler:               controller.add,
	}

	return []rest.Route{r1, r2}
}

func (controller *deviceModelController) getAll(ctx *rest.RequestContext) (interface{}, rest.Error) {
	return controller.deviceModelService.GetAll(), nil
}

func (controller *deviceModelController) add(ctx *rest.RequestContext) (interface{}, rest.Error) {
	model := models.CarModelModel{}

	ctx.DecodeBody(model)
	err := controller.deviceModelService.Add(model)

	if err != nil {
		return nil, rest.NewRestError(http.StatusBadRequest)
	}

	return rest.NewSuccess("devicemodel.add"), nil
}