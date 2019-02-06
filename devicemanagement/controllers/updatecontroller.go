package controllers

import (
	"net/http"

	"nova/devicemanagement/models"
	"nova/devicemanagement/services"
	"nova/rest"
)

type updateController struct {
	UpdateService services.UpdateService
}

func NewUpdateController(updateService services.UpdateService) *updateController {
	return &updateController{updateService}
}

func (controller *updateController) GetRoutes() []rest.Route {
	r1 := rest.Route{
		Method:                rest.MethodPost,
		Path:                  "/updates",
		RequiresAuthorization: true,
		Handler:               controller.CreateUpdate,
	}

	r2 := rest.Route{
		Method:                rest.MethodGet,
		Path:                  "/updates",
		RequiresAuthorization: true,
		Handler:               controller.GetAll,
	}

	r3 := rest.Route{
		Method:                rest.MethodGet,
		Path:                  "/updates/{id}",
		RequiresAuthorization: true,
		Handler:               controller.GetAll,
	}

	return []rest.Route{r1, r2, r3}
}

func (controller *updateController) GetAll(ctx *rest.RequestContext) (interface{}, rest.Error) {
	updates := controller.UpdateService.GetAllUpdates()
	return updates, nil
}

func (controller *updateController) GetAllUpdateTasks(ctx *rest.RequestContext) (interface{}, rest.Error) {
	updateTasks := controller.UpdateService.GetAllUpdateTasks()
	return updateTasks, nil
}

func (controller *updateController) GetById(ctx *rest.RequestContext) (interface{}, rest.Error) {
	//TODO: Implement
	_, ok := ctx.Vars["id"]

	if !ok {
		return nil, rest.NewRestError(http.StatusNotFound)
	}

	return nil, nil
}

func (controller *updateController) CreateUpdate(ctx *rest.RequestContext) (interface{}, rest.Error) {
	updateModel := &models.UpdateModel{}

	err := ctx.DecodeBody(&updateModel)
	if err != nil {
		return nil, err
	}
	err1 := controller.UpdateService.CreateUpdate(*updateModel)

	if err1 != nil {
		return nil, rest.NewRestError(http.StatusBadRequest)
	}
	return rest.NewSuccess("CreateUpdate"), nil
}

func (controller *updateController) AssignUpdateToCar(ctx *rest.RequestContext) (interface {}, rest.Error) {
	updateTaskModel := &models.UpdateTaskModel{}

	err := ctx.DecodeBody(&updateTaskModel)

	if err != nil {
		return nil, rest.NewRestError(http.StatusBadRequest)
	}

	err1 := controller.UpdateService.UpdateDevice(*updateTaskModel)

	if err1 != nil {
		return nil, rest.NewRestError(http.StatusBadRequest)
	}

	return rest.NewSuccess("AssignUpdateToCar"), nil

}