package validation

import (
	"github.com/adel-hadadi/translator/config"
	"github.com/adel-hadadi/translator/internal/api/translate"
	"github.com/go-playground/validator/v10"
)

type Validation struct {
	Cfg        *config.Config
	Validator  *validator.Validate
	Translator *translate.Translator
}

func NewValidator(cfg *config.Config, tr *translate.Translator) *Validation {
	return &Validation{
		Cfg:        cfg,
		Validator:  validator.New(),
		Translator: tr,
	}
}

func (v *Validation) Validate(i interface{}) []translate.ErrValidation {
	if err := v.Validator.Struct(i); err != nil {
		var validationErrs []translate.ErrValidation
		for _, fieldError := range err.(validator.ValidationErrors) {
			validationErrs = append(
				validationErrs,
				v.Translator.TranslateValidationError(fieldError),
			)
		}

		return validationErrs
	}

	return nil
}
