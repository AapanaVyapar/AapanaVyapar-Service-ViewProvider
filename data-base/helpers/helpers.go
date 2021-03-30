package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Validate(s interface{}) error {

	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		myErr := ""
		for _, err := range err.(validator.ValidationErrors) {
			myErr += " " + err.Field()
		}

		return fmt.Errorf("Error in validating :" + myErr)
	}

	return nil
}
