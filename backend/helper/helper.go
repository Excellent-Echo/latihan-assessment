package helper

import (
	"errors"
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func SplitErrorInformation(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func ValidateIDNumber(id string) error {
	if num, err := strconv.Atoi(id); err != nil || num == 0 || math.Signbit(float64(num)) == true {
		return errors.New("input must be a valid id")
	}

	return nil
}
