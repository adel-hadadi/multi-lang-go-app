package routes

import (
	"github.com/adel-hadadi/translator/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewAuthRoutes(router *gin.RouterGroup, handler handlers.AuthHandler) {
	router.POST("/signup", handler.SignUp)
}
