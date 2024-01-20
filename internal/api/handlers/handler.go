package handlers

import (
	"github.com/adel-hadadi/translator/internal/api/services"
	"github.com/adel-hadadi/translator/internal/api/validation"
)

type Handlers struct {
	AuthHandler *AuthHandler
}

func NewHandlers(service *services.Services, validator *validation.Validation) *Handlers {
	return &Handlers{
		AuthHandler: NewAuthHandler(service.AuthService, validator),
	}
}
