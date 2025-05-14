package setup

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func SetupLogger(app *echo.Echo) {
	app.Use(middleware.RequestID())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} [MIDDLEWARE] (${id}) " +
			"- ${method} ${uri} -> HTTP ${status} (${latency_human}) ${error}",
	}))

	app.Logger.SetHeader("${time_rfc3339_nano} [${level}] (${id}) ${short_file}:${line} =>")
	app.Logger.SetLevel(log.INFO)
}
