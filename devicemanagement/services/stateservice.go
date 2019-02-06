package services

import (
	"net/http"
)

type StateService interface {
	RestoreFromRequest(r *http.Request) (*UserToken, error)
	SetToResponse(r *http.Request, w *http.ResponseWriter, token UserToken) error
}