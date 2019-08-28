package controller


import (
	"net/http"
	"github.com/labstack/echo"
)

func (c *Core) Login(ctx echo.Context) error {
	type request struct {
		// in: body
		Body struct {
			// Required: true
			Nickname string `json:"nickname" validate:"required"`
			// Required: true
			Password string `json:"password" validate:"required"`
		}
	}

	// Request validation.
	req := new(request)
	if err := ctx.Bind(req); err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, []string{err.Error()})
	} else if err = c.validator.Validate(req); err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, []string{err.Error()})
	}


	return c.response.OK(ctx, http.StatusOK, nil)
}