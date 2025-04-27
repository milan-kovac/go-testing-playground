package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *UserService
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Register endpoint is not implemented yet",
	})

}
