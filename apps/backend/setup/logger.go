package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupLogger(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} [MIDDLEWARE] (${pid} - ${id}) " +
			"- ${method} ${path} -> HTTP ${status} (${latency}) ${error}\n",
	}))
	log.SetLevel(log.LevelInfo)
	// app.Use(middleware.RequestID())
	// app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "${time_rfc3339_nano} [MIDDLEWARE] (${id}) " +
	// 		"- ${method} ${uri} -> HTTP ${status} (${latency_human}) ${error}\n",
	// }))

	// app.Logger.SetHeader("${time_rfc3339_nano} [${level}] (${id}) ${short_file}:${line} =>")
	// app.Logger.SetLevel(log.INFO)
}
