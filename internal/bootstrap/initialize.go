package bootstrap

import (
	"fmt"

	"github.com/ahmdyaasiin/workshop-ci-cd/internal/infra/config"
	"github.com/ahmdyaasiin/workshop-ci-cd/internal/infra/db"
	"github.com/ahmdyaasiin/workshop-ci-cd/internal/infra/http"
	logger "github.com/ahmdyaasiin/workshop-ci-cd/internal/infra/log"
	"github.com/gofiber/fiber/v2"

	ph "github.com/ahmdyaasiin/workshop-ci-cd/internal/app/product/handler"
	pr "github.com/ahmdyaasiin/workshop-ci-cd/internal/app/product/repository"
	pu "github.com/ahmdyaasiin/workshop-ci-cd/internal/app/product/usecase"
)

type App struct {
	config *config.Env
	db     db.Postgres
	server *fiber.App
}

func Initialize() error {
	_config, err := config.Load()
	if err != nil {
		return err
	}

	_db, err := db.NewPostgres(_config.DB)
	if err != nil {
		return err
	}

	if err := logger.NewZeroLog(); err != nil {
		return err
	}

	app := &App{
		config: _config,
		db:     _db,
		server: http.NewFiber(),
	}

	if err := app.handleFlags(); err != nil {
		return err
	}

	return app.start()
}

func (app *App) start() error {
	app.health()
	app.registerRoutes()

	return app.server.Listen(fmt.Sprintf(":%d", app.config.AppPort))
}

func (app *App) health() {
	app.server.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}

func (app *App) registerRoutes() {
	productRepository := pr.New(app.db.GetConnection())
	productUseCase := pu.New(productRepository)
	productHandler := ph.New(productUseCase)

	router := app.server.Group("")
	productHandler.MountRoutes(router)
}
