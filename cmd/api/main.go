package main

import (
	"github.com/adel-hadadi/translator/cmd"
	"github.com/adel-hadadi/translator/config"
	"github.com/adel-hadadi/translator/internal/api"
)

func main() {
	cfg := config.NewConfig()

	app := cmd.NewApplication(cfg)

	api.InitServer(app.Cfg, app.Handlers)
}
