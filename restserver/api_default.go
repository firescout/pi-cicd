package restserver

import (
	"net/http"
	"strings"
)

type DefaultApiController struct {
	service      DefaultApiServicer
	errorHandler ErrorHandler
}

type DefaultApiOption func(*DefaultApiController)

func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"OnPush",
			strings.ToUpper("Get"),
			"/repo/push",
			c.OnPush,
		},
		{
			"GetShutdown",
			strings.ToUpper("Get"),
			"/shutdown",
			c.GetShutdown,
		},
	}
}

func (c *DefaultApiController) OnPush(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.OnPush(r.Context(), r.URL.Query().Get("repo"))
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	EncodeJSONResponse(result.Body, &result.Code, w)
}

func (c *DefaultApiController) GetShutdown(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetShutdown(r.Context())
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	EncodeJSONResponse(result.Body, &result.Code, w)

}
