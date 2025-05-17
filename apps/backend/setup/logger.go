package setup

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger(app *fiber.App) {
	logger := log.Output(zerolog.ConsoleWriter{
		Out:         os.Stdout,
		TimeFormat:  time.TimeOnly,
		FieldsOrder: []string{"requestId"},
	})
	log.Logger = logger
	app.Use(requestid.New())
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
		Fields: []string{"pid", "requestId", "method", "url", "status", "latency", "error"},
	}))
	app.Use(func(c *fiber.Ctx) error {
		reqId := c.Locals(requestid.ConfigDefault.ContextKey).(string)
		ctx := context.WithValue(c.UserContext(), "requestId", reqId)
		reqLogger := logger.With().Str("requestId", reqId).Logger()
		ctx = reqLogger.WithContext(ctx)
		c.SetUserContext(ctx)
		return c.Next()
	})
}
