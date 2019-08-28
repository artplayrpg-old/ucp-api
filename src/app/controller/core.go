package controller

import (
	"app"
	"app/config"
)

// Core represents the core for all the handlers.
type Core struct {
	conf		*config.Config
	validator	app.IValidate
	response	app.IResponse
}

// New returns a core for all the handlers.
func New(conf *config.Config, validator app.IValidate, resp app.IResponse,) *Core {
	return &Core{
		conf: conf,
		validator: validator,
		response: resp,
	}
}