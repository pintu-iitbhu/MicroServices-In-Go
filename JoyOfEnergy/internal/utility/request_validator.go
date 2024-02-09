package utility

import (
	"JoyOfEnergy/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

type Validator struct{}

func NewValidator() *validator.Validate {
	requestValidator := validator.New()
	log := logger.NewLogger()
	requestValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	var err error
	defer func() {
		if err != nil {
			panic("Validator registration failed " + err.Error())
		}
	}()

	err = requestValidator.RegisterValidation("notNilUUID", NotNilUUIDStringValidator)
	if err != nil {
		log.Error("Error while registering custom validator func NotNilUUIDValidator %s\n", err.Error())
		return nil
	}
	return requestValidator
}

func NotNilUUIDStringValidator(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	uuidValue, err := uuid.Parse(fieldValue)
	if err != nil || uuidValue == uuid.Nil {
		return false
	}
	return true
}
