package helper

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator2 "github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strings"
)

type ValidateErrorMessage struct {
	Message []string `json:"validation_errors"`
}

func DtoValidate(dto interface{}) (validationResult *[]byte) {
	validator := validator2.New()
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validator.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator2.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := validator.Struct(dto)
	var msg []string
	if err != nil {
		for _, err := range err.(validator2.ValidationErrors) {
			msg = append(msg, err.Translate(trans))
		}
		body := ValidateErrorMessage{
			Message: msg,
		}
		var json = jsoniter.ConfigCompatibleWithStandardLibrary

		serialisedMsg, _ := json.Marshal(&body)
		return &serialisedMsg
	}
	return nil
}

func MaptwoStructs[T1 any, T2 any](source *T1, target *T2) error {
	err := "Marshal JSON crash"
	byteArray, err1 := jsoniter.Marshal(source)
	if err1 != nil {
		return errors.New(err)
	}
	err2 := jsoniter.Unmarshal(byteArray, target)
	if err2 != nil {
		return errors.New(err)
	}
	return nil
}

func MapStructToJSON[T any](structIn *any) string {
	jsonStr, _ := jsoniter.MarshalToString(structIn)
	return jsonStr
}
