package controllers

import (
	"net/http"
	"strconv"

	"nova/devicemanagement/models"
	"nova/devicemanagement/services"
	"nova/rest"
)

type updateTaskController struct {
	UpdateService services.UpdateService
}

func NewUpdateTaskController(updateService services.UpdateService) *updateTaskController {
	return &updateTaskController{updateService}
}

func (controller *updateTaskController) GetRoutes() []rest.Route {
	r1 := rest.Route{
		Method:                rest.MethodGet,
		Path:                  "/updates/tasks/",
		RequiresAuthorization: true,
		Handler:               controller.GetAll,
	}

	r2 := rest.Route{
		Method:                rest.MethodPost,
		Path:                  "/updates/tasks",
		RequiresAuthorization: true,
		Handler:               controller.Add,
	}

	r3 := rest.Route{
		Method:                rest.MethodGet,
		Path:                  "/updates/tasks/{id}",
		RequiresAuthorization: true,
		Handler:               controller.GetById,
	}

	return []rest.Route{r1, r2, r3}
}

func (controller *updateTaskController) GetAll(ctx *rest.RequestContext) (interface{}, rest.Error) {
	return controller.UpdateService.GetAllUpdateTasks(), nil
}

func (controller *updateTaskController) GetById(ctx *rest.RequestContext) (interface{}, rest.Error) {

	strId, ok := ctx.Vars["id"]

	if !ok {
		return nil, rest.NewRestError(http.StatusNotFound)
	}

	id, err := strconv.Atoi(strId)

	if err != nil {
		return nil, rest.NewRestError(http.StatusBadRequest)
	}

	return controller.UpdateService.GetUpdateTaskById(uint(id)), nil
}

func (controller *updateTaskController) Add(ctx *rest.RequestContext) (interface{}, rest.Error) {
	updateTaskModel := &models.UpdateTaskModel{}

	err := ctx.DecodeBody(updateTaskModel)

	if err != nil {
		return nil, rest.NewRestError(http.StatusBadRequest)
	}

	err1 := controller.UpdateService.UpdateDevice(*updateTaskModel)

	if err1 != nil {
		return nil, rest.NewRestError(http.StatusInternalServerError)
	}

	return rest.NewSuccess("Add"), nil
}