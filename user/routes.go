package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/milan-kovac/database"
	"github.com/milan-kovac/middlewares"
)

func UserRoutes(app *fiber.App) {
	userRepository := NewUserRepository(database.DB)
	userService := NewUserService(userRepository)
	userController := NewUserController(userService)

	userGroup := app.Group("/users")

	userGroup.Post(
		"/",
		middlewares.ValidateBody[CreateUserRequest](),
		userController.Create,
	)

	userGroup.Get(
		"/",
		userController.GetAll,
	)

	userGroup.Get(
		"/:id",
		middlewares.ValidateIdParam(),
		userController.Get,
	)
}
