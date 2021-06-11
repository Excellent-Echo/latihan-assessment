package helper

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"math"
	"strconv"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(msg string, code int, status string, data interface{}) Response {
	var meta = Meta{
		Message: msg,
		Code:    code,
		Status:  status,
	}

	var response = Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func ErrorInformation(getError error) []string {
	var errortest []string

	for _, e := range getError.(validator.ValidationErrors) {
		errortest = append(errortest, e.Error())
	}

	return errortest
}

func ValidateIDNumber(ID string) error {
	if number, err := strconv.Atoi(ID); err != nil || number == 0 || math.Signbit(float64(number)) == true {
		return errors.New("your input error, id user must be valid")
	}
	return nil
}
