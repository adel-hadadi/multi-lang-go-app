package handlers

import (
	"github.com/adel-hadadi/translator/internal/api/dto"
	"github.com/adel-hadadi/translator/internal/api/services"
	"github.com/adel-hadadi/translator/internal/api/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	AuthService *services.AuthService
	Validator   *validation.Validation
}

func NewAuthHandler(authService *services.AuthService, validator *validation.Validation) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Validator:   validator,
	}
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	req := dto.SignupReq{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.Validator.Validate(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	_, err := h.AuthService.Signup(req, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
}
