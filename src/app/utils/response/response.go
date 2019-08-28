package response

import "github.com/labstack/echo"

type OKResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int      `json:"status"`
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

// Output is the response object.
type Output struct{}

// New returns a new response object.
func New() *Output {
	return &Output{}
}

// OK will write an OK status.
func (o *Output) OK(ctx echo.Context, status int, data interface{}) error {
	return ctx.JSON(status, OKResponse{
		status,
		true,
		data,
	})
}

// Error will write an Error status.
func (o *Output) Error(ctx echo.Context, status int, errors []string) error {
	return ctx.JSON(status, ErrorResponse{
		status,
		false,
		errors,
	})
}