// Novo Device Updater
// Felix Almesberger 2018 Pentasys AG

package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"nova/devicemanagement/services"
)

// statics are not so nice
var StateService services.StateService

type Route struct {
	Path                  string
	RequiresAuthorization bool
	Method                string
	Handler               func(ctx *RequestContext) (interface{}, Error)
}

func (route Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if route.RequiresAuthorization && !requestIsAuthorized(r) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//prepare request context
	vars := mux.Vars(r)
	ctx := &RequestContext{Request: r, ResponseWriter: w, Vars: vars}

	result, executionError := route.Handler(ctx)
	if executionError != nil {
		w.WriteHeader(executionError.Status())
		return
	}

	//deliver success response if action has no return value
	if result == nil && executionError == nil {
		result = NewSuccess(route.Path)
	}

	encoder := json.NewEncoder(w)
	encodingErr := encoder.Encode(result)

	if encodingErr == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		log.Fatalf("Could not encode response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func requestIsAuthorized(r *http.Request) bool {
	token, err := StateService.RestoreFromRequest(r)

	if err != nil {
		return false
	}

	return token.Authenticated
}