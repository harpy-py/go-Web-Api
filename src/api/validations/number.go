package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func IranianMobileNumberValidator(fld validator.FieldLevel) bool{
	value, ok := fld.Field().Interface().(string)
	if !ok{
		return false
	}

	res, err := regexp.MatchString(`^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`, value)
	if err != nil {
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return res
}