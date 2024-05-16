package model

import "github.com/go-playground/validator/v10"

type PostInput struct {
	Title   string   `json:"title" form:"title" validate:"required"`
	Content string   `json:"content" form:"content" validate:"required"`
	Tags    []string `json:"tags" form:"tags" validate:"required"`
}

func (input *PostInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
