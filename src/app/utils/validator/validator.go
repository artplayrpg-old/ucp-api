package validator

import (
	v9 "gopkg.in/go-playground/validator.v9"
)

type Validate struct {
	validator	*v9.Validate
}

func New() *Validate {
	return &Validate{
		validator: v9.New(),
	}
}

func (v *Validate) Validate(data interface{}) (err error) {
	return v9.New().Struct(data)
}