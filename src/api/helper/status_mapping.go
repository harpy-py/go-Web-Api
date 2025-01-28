package helper

import (
	"net/http"

	"github.com/harpy-py/go-Web-Api/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	service_errors.OtpExists: 409,
	service_errors.OtpUsed: 409,
	service_errors.OtpInvalid: 400,
}

func TranslateErrorToStatusCode(err error) int{
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}