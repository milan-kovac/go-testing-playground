package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/milan-kovac/config"
	"github.com/milan-kovac/database"
)

func main() {
	app := fiber.New()

	// APP CONFIG
	config.LoadConfig()

	// DATABASE CONFIG
	database.Connect()

	app.Listen(":" + config.Env.Port)

}
