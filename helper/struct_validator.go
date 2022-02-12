package helper

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
	"lienquanMess/log"
	"strings"
)

type StructValidator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     ut.Translator
}

func NewStructValidaten() *StructValidator {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")

	return &StructValidator{
		Validator: validator.New(),
		Uni:       uni,
		Trans:     trans,
	}
}

func (cv *StructValidator) RegisterValidate() {
	if err := en_translations.RegisterDefaultTranslations(cv.Validator, cv.Trans); err != nil {
		log.Error(err.Error())
	}

	cv.Validator.RegisterValidation("pwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 8
	})

	cv.Validator.RegisterTranslation("required", cv.Trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} là bắt buộc", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	cv.Validator.RegisterTranslation("e164", cv.Trans, func(ut ut.Translator) error {
		return ut.Add("e164", "{0} không hợp lệ", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("e164", fe.Field())
		return t
	})

	cv.Validator.RegisterTranslation("pwd", cv.Trans, func(ut ut.Translator) error {
		return ut.Add("pwd", "Mật khẩu tối thiểu 8 kí tự", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("pwd", fe.Field())
		return t
	})
}

func (cv *StructValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err == nil {
		return nil
	}

	transErrors := make([]string, 0)
	for _, e := range err.(validator.ValidationErrors) {
		transErrors = append(transErrors, e.Translate(cv.Trans))
	}
	return errors.Errorf("%s", strings.Join(transErrors, " \n "))
}
