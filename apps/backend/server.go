package main

import (
	"context"
	"net/http"
	"tsm/domain"
	"tsm/domain/auth"
	"tsm/setup"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()
	domain.SetupFlags()

	setup.SetupLogger(app)

	pool, err := domain.NewDatabasePool(context.Background())

	if err != nil {
		log.Fatal().Stack().Err(err)
	}

	defer pool.Close()

	if err = Seed(context.Background(), pool); err != nil {
		log.Fatal().Stack().Err(err)
	}

	app.Mount("/auth", auth.Routes(pool)).Name("Authentication routes")

	if err := app.Listen(":4000"); err != nil && err != http.ErrServerClosed {
		log.Fatal().Stack().Err(err)
	}
}
