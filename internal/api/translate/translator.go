package translate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
	"strings"
)

type Translator struct {
	ValidationErrs map[string]string
	Fields         map[string]string
	Messages       map[string]interface{}
}

type ErrValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewTranslator(local string) (*Translator, error) {
	validationErr, messages, err := readLangFiles(local)
	if err != nil {
		return nil, err
	}

	return &Translator{
		ValidationErrs: validationErr["validations"],
		Fields:         validationErr["fields"],
		Messages:       messages,
	}, nil
}

func readLangFiles(local string) (map[string]map[string]string, map[string]interface{}, error) {
	var validationErr map[string]map[string]string
	var messages map[string]interface{}

	file, err := os.ReadFile(fmt.Sprintf("lang/%s/validation.json", local))
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("error occurred reading validation file: %s", err))
	}

	err = json.Unmarshal(file, &validationErr)
	if err != nil {
		return nil, nil, err
	}

	messageFile, err := os.ReadFile(fmt.Sprintf("lang/%s/validation.json", local))
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("error occurred reading message file: %s", err))
	}

	err = json.Unmarshal(messageFile, &messages)
	if err != nil {
		return nil, nil, err
	}

	return validationErr, messages, nil
}

func (t *Translator) TranslateValidationError(err validator.FieldError) ErrValidation {
	field, has := t.Fields[err.Field()]
	if !has {
		field = strings.ToLower(err.Field())
	}
	message, has := t.ValidationErrs[err.Tag()]
	if !has {
		message = err.Tag()
	}

	if err.Param() == "" {
		return ErrValidation{
			Field: strings.ToLower(err.Field()),
			Message: fmt.Sprintf(
				message,
				field,
			),
		}
	}

	return ErrValidation{
		Field: strings.ToLower(err.Field()),
		Message: fmt.Sprintf(
			message,
			field,
			err.Param(),
		),
	}

}

func (t *Translator) Message(key string) string {
	messages := t.Messages
	keys := strings.Split(key, ".")

	for _, k := range keys {
		if nested, ok := messages[k].(map[string]interface{}); ok {
			messages = nested
		} else {
			if strVal, ok := messages[k].(string); ok {
				return strVal
			}
			return ""
		}
	}

	return ""
}
