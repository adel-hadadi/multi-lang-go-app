package services

import (
	"github.com/adel-hadadi/translator/internal/api/translate"
)

type Services struct {
	AuthService *AuthService
}

func NewServices(tr *translate.Translator) *Services {
	return &Services{
		AuthService: NewAuthService(tr),
	}
}
