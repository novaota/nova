package rest

import (
	"encoding/json"
	"net/http"
)

type RequestContext struct {
	Vars map[string]string
	ResponseWriter http.ResponseWriter
	Request *http.Request
}

func (ctx *RequestContext) DecodeBody(out interface {}) Error {
	decoder := json.NewDecoder(ctx.Request.Body)

	//wrap normal error in http error
	if err := decoder.Decode(&out); err != nil {
		out = nil
		return NewRestError(http.StatusBadRequest)
	}
	return nil
}
