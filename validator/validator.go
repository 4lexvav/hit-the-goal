package validator

import (
	"net/http"
	"reflect"
	"strings"
	"sync"

	exterrors "github.com/eugeneradionov/ext-errors"
	v "github.com/go-playground/validator/v10"
)

var (
	validatorInstance *v.Validate
	once              = &sync.Once{}
)

func Get() *v.Validate {
	return validatorInstance
}

func FormatErrors(errs v.ValidationErrors) exterrors.ExtErrors {
	var validationErrs = &exterrors.Errors{
		Errs: make([]exterrors.ExtError, 0, len(errs)),
	}

	for i := range errs {
		validationErrs.Add(exterrors.Error{
			Field:       errs[i].Field(),
			Description: errs[i].Tag(),
			Message:     errs[i].Error(),
			Code:        http.StatusUnprocessableEntity,
		})
	}

	return validationErrs
}

func Load() (err error) {
	once.Do(func() {
		validatorInstance = v.New()
		validatorInstance.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}

			return name
		})
	})

	return err
}
