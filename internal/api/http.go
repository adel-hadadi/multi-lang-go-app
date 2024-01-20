package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/adel-hadadi/translator/config"
	"github.com/adel-hadadi/translator/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitServer(cfg *config.Config, handler *handlers.Handlers) {
	router := gin.Default()
	initRoutes(router, handler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.App.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
