package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"bun-spreader/config"
	"bun-spreader/routes"
	"bun-spreader/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${url} ${error} ${time}\n",
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, err interface{}) {
			log.Printf("Panic Recovered: %v", err)
		},
	}))
	err := config.Init()
	utils.HandleError(err)

	routes.RegisterRoutes(app)
	app.Listen(":9000")
}
