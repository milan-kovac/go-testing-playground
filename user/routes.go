package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/milan-kovac/middlewares"
)

func UserRoutes(app *fiber.App) {
	userController := &UserController{}

	userGroup := app.Group("/users")

	userGroup.Post(
		"/",
		middlewares.ValidateBody[CreateTaskRequest](),
		userController.Create,
	)
}
