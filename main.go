package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/milan-kovac/config"
	"github.com/milan-kovac/database"
	"github.com/milan-kovac/user"
)

func main() {
	app := fiber.New()

	// APP CONFIG
	config.LoadConfig()

	// DATABASE CONFIG
	database.Connect()

	// REGISTER ROUTES
	user.UserRoutes(app)

	app.Listen(":" + config.Env.Port)

}
