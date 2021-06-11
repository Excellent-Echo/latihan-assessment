package helper

import (
	"errors"
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func SpilitErrorInformation(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func ValidateIDNumber(ID string) error {
	if num, err := strconv.Atoi(ID); err != nil || num == 0 || math.Signbit(float64(num)) == true {
		return errors.New("Input must be valid")
	}

	return nil
}
