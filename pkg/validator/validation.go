package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type IError struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

/*
Validate: validates the fields of the incoming request
and returns the list of errors indicating bad fields for users
*/
func Validate(data interface{}) []*IError {

	// check for empty field
	var errors []*IError

	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)

	//Use reflector to reverse engineer struct
	reflected := reflect.ValueOf(data)

	if err != nil {

		// if validation syntax is invalid for bad passed value
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		//iterate over the validation errors
		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, ok := reflected.Type().FieldByName(err.StructField())
			if !ok {
				fmt.Println("Error occured while fetching fields from struct")
			}

			var name string
			// if json tag doesn't exist use lowe case name of struct field
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			var value string

			switch err.Tag() {
			case "required":
				value = "The " + name + " is required"
			case "email":
				value = "The " + name + " should be a valid email"
			case "eqfield":
				value = "The " + name + " should be equal to the " + err.Param()
			case "max":
				value = "The " + name + " should be max length of " + err.Param()
			case "min":
				value = "The " + name + " should be minimum length of " + err.Param()
			case "numeric", "boolean":
				value = "The " + name + " should be of type " + err.Tag()
			default:
				value = "The " + name + " is invalid"
			}

			errors = append(errors, &IError{
				Field: name,
				Value: value,
			})

		}
		return errors
	}
	return nil
}
