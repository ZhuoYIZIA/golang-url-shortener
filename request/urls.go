package request

import (
	"github.com/go-playground/validator/v10"
)

type UrlRequest struct {
	Url string `json:"url" validate:"required,url"`
}

var validate *validator.Validate

func (url UrlRequest) Validator() error {
	validate = validator.New()
	err := validate.Struct(url)
	if err != nil {
		return err
	}
	return nil
}
