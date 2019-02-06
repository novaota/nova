package controllers

import (
	"net/http"

	"nova/devicemanagement/models"
	"nova/devicemanagement/services"
	"nova/rest"
)

type loginController struct {
	services.AuthenticationService
	services.StateService
}

func NewLoginController(authService services.AuthenticationService, stateService services.StateService) *loginController {
	return &loginController{authService, stateService}
}

func (controller *loginController) GetRoutes() []rest.Route {
	r1 := &rest.Route{
		Method:                rest.MethodPost,
		Path:                  "/login",
		RequiresAuthorization: false,
		Handler:               controller.handleLogin,
	}

	return []rest.Route{*r1}
}

func (controller *loginController) handleLogin(ctx *rest.RequestContext) (interface{}, rest.Error) {
	user := &models.UserModel{}

	err := ctx.DecodeBody(&user)

	if err != nil {
		return nil, rest.NewRestError(http.StatusBadRequest)
	}

	authService := controller.AuthenticationService
	isRegistered := authService.IsRegistered(user.Username, user.Password)

	if !isRegistered {
		return nil, rest.NewRestError(http.StatusUnauthorized)
	}

	token := services.UserToken{
		Username:      user.Username,
		Authenticated: true,
	}
	controller.StateService.SetToResponse(ctx.Request, &ctx.ResponseWriter, token)

	return rest.NewSuccess("login"), nil
}
