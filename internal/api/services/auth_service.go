package services

import (
	"errors"
	"github.com/adel-hadadi/translator/internal/api/dto"
	"github.com/adel-hadadi/translator/internal/api/translate"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	Translator *translate.Translator
}

func NewAuthService(tr *translate.Translator) *AuthService {
	return &AuthService{
		Translator: tr,
	}
}

func (s *AuthService) Signup(req dto.SignupReq, ctx *gin.Context) (*dto.SignupRes, error) {
	if req.Username == "admin" {
		return nil, errors.New(s.Translator.Message("users.already_exists"))
	}

	return nil, nil
}
