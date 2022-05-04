package request

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type GetUrlRequest struct {
	Id string `json:"id" validate:"required,alphanum"`
}

func (request GetUrlRequest) Validator() error {
	validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

type CreateUrlRequest struct {
	Url string `json:"url" validate:"required,url"`
}

func (request CreateUrlRequest) Validator() error {
	validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}
