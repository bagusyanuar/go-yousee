package common

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type (
	CustomValidator struct {
		Validator *validator.Validate
	}

	ErrorValidation struct {
		Field string
		Tag   string
		Value interface{}
	}
)

func (v CustomValidator) Validate(data interface{}) []string {
	validationErrors := []string{}
	errs := v.Validator.Struct(data)
	if errs != nil {
		localeID := id.New()
		uni := ut.New(localeID, localeID)
		trans, _ := uni.GetTranslator("id")
		id_translations.RegisterDefaultTranslations(v.Validator, trans)
		for _, err := range errs.(validator.ValidationErrors) {
			// var element ErrorValidation
			// element.Field = err.Field()
			// element.Tag = err.Error()
			// element.Value = err.Value()
			// validationErrors = append(validationErrors, element)
			validationErrors = append(validationErrors, strings.ToLower(err.Translate(trans)))
			fmt.Println(err.Translate(trans))
		}
	}
	return validationErrors
}

func (v CustomValidator) TranslateError(errs []ErrorValidation) []string {

	var errMessages []string
	for _, err := range errs {
		errMessages = append(errMessages, fmt.Sprintf("%s : need to implement %s",
			err.Field,
			err.Tag,
		))
	}
	return errMessages
}
