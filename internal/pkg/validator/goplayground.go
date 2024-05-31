package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	goPlaygroundValidator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type GoPlaygroundValidator struct {
	validator   *goPlaygroundValidator.Validate
	translation ut.Translator
}

func NewGoPlayValidator() *GoPlaygroundValidator {
	vl := GoPlaygroundValidator{}

	gplv := goPlaygroundValidator.New()
	gplv.RegisterTagNameFunc(vl.tagName)

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	vl.validator = gplv
	vl.translation = trans

	// register default translation
	enTranslations.RegisterDefaultTranslations(vl.validator, trans)

	return &vl
}

func (v *GoPlaygroundValidator) ValidateStruct(data interface{}) (map[string]string, error) {
	if err := v.validator.Struct(data); err != nil {
		return v.customErrorMessage(err), err
	}

	return nil, nil
}

func (v *GoPlaygroundValidator) tagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func (m *GoPlaygroundValidator) customErrorMessage(err error) map[string]string {
	errMap := make(map[string]string)
	errs, ok := err.(goPlaygroundValidator.ValidationErrors)
	if !ok {
		return errMap
	}

	regexKey := regexp.MustCompile(`^[^.]*.`)
	regexDate := regexp.MustCompile("2006-01-02")
	regexTime := regexp.MustCompile("15:04:05")
	for _, e := range errs {
		field := regexKey.ReplaceAllString(e.Namespace(), "")

		value := e.Translate(m.translation)
		value = regexp.MustCompile(e.Field()).ReplaceAllString(value, field)

		// replace message
		tag := e.Tag()
		fmt.Println(tag)
		switch e.Tag() {
		case "date":
			value = regexDate.ReplaceAllString(value, "yyyy-MM-dd")
		case "datetime":
			value = regexDate.ReplaceAllString(value, "yyyy-MM-dd")
			value = regexTime.ReplaceAllString(value, "HH:mm:ss")
		case "required_with":
			value = field + " is a required field"
		}

		errMap[field] = value
	}

	return errMap
}
