package controllers

import (
	"nova/devicemanagement/models"
	"nova/devicemanagement/services"
	"nova/rest"
)

type ownerController struct {
	ownerService services.OwnerService
}

func NewOwnerController(ownerService services.OwnerService) *ownerController {
	return &ownerController{ownerService}
}

func (controller *ownerController) GetRoutes() []rest.Route {
	r1 := &rest.Route{
		Method:                rest.MethodGet,
		Path:                  "/owners",
		RequiresAuthorization: true,
		Handler:               controller.GetAllOwners,
	}

	r2 := &rest.Route{
		Method:                rest.MethodPost,
		Path:                  "/owners",
		RequiresAuthorization: true,
		Handler:               controller.AddOwner,
	}

	return []rest.Route{*r1, *r2}
}

func (controller *ownerController) GetAllOwners(ctx *rest.RequestContext) (interface{}, rest.Error) {
	return controller.ownerService.GetAll(), nil
}

func (controller *ownerController) AddOwner(ctx *rest.RequestContext) (interface{}, rest.Error) {
	model := models.OwnerModel{}

	err := ctx.DecodeBody(&model)

	controller.ownerService.Add(model)

	return rest.NewSuccess("Owners.Add"), err
}
