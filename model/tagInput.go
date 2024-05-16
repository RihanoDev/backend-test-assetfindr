package model

import "github.com/go-playground/validator/v10"

type TagInput struct {
	Label string `json:"label" form:"label"`
}

func (input *TagInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
