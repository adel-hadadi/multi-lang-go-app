package api

import (
	"github.com/adel-hadadi/translator/internal/api/handlers"
	"github.com/adel-hadadi/translator/internal/api/routes"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine, handlerContainer *handlers.Handlers) {
	goals := router.Group("/goals")
	{
		routes.InitGoalRoutes(goals)
	}

	auth := router.Group("/auth")
	{
		routes.NewAuthRoutes(auth, *handlerContainer.AuthHandler)
	}
}
