package exceptions

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	switch validationError.Tag() {
	case "required":
		return errors.New(validationError.StructField() + " é obrigatório")
	case "max":
		return errors.New(validationError.StructField() + " deve conter no máximo " + validationError.Param() + " caracteres")
	case "min":
		return errors.New(validationError.StructField() + " deve conter no mínimo " + validationError.Param() + " caracteres")
	case "email":
		return errors.New(validationError.StructField() + " deve ser válido")
	}
	return nil
}
