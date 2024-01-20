package cmd

import (
	"github.com/adel-hadadi/translator/config"
	"github.com/adel-hadadi/translator/internal/api/handlers"
	"github.com/adel-hadadi/translator/internal/api/services"
	"github.com/adel-hadadi/translator/internal/api/translate"
	"github.com/adel-hadadi/translator/internal/api/validation"
)

type Application struct {
	Cfg       *config.Config
	Validator *validation.Validation
	Services  *services.Services
	Handlers  *handlers.Handlers
}

func NewApplication(cfg *config.Config) *Application {
	tr, err := translate.NewTranslator(cfg.App.Local)
	if err != nil {
		panic(err.Error())
	}
	validator := validation.NewValidator(cfg, tr)
	service := services.NewServices(tr)
	handler := handlers.NewHandlers(service, validator)

	return &Application{
		Cfg:       cfg,
		Validator: validator,
		Services:  service,
		Handlers:  handler,
	}
}
