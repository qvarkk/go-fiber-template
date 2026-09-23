package config

import (
	"fmt"
	"slices"

	"github.com/go-playground/validator/v10"
)

func newValidator() (*validator.Validate, error) {
	validate := validator.New()

	err := validate.RegisterValidation("validEnv", func(fl validator.FieldLevel) bool {
		input := AppEnv(fl.Field().String())
		return slices.Contains(AllowedEnvironments(), input)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register validEnv validator: %w", err)
	}

	err = validate.RegisterValidation("validLogLevel", func(fl validator.FieldLevel) bool {
		input := LogLevel(fl.Field().String())
		return slices.Contains(AllowedLogLevels(), input)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register validLogLevel validator: %w", err)
	}

	return validate, nil
}
