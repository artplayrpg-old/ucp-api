package app

import "github.com/labstack/echo"

// IBind provides bind and validation for requests.
type IValidate interface {
	Validate(v interface{}) error
}

// IResponse provides outputs for data.
type IResponse interface {
	OK(ctx echo.Context, status int, data interface{}) error
	Error(ctx echo.Context, status int, errors []string) error
}