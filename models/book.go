package models

import (
    "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}

type Book struct {
    ID     uint   `json:"id"`
    Title  string `json:"title" validate:"required"`
    Author string `json:"author" validate:"required"`
}

// ValidateBook validates a book struct
func ValidateBook(book *Book) error {
    return validate.Struct(book)
}
