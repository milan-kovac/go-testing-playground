package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ValidateIdParam() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		idStr := ctx.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID.",
			})
		}

		ctx.Locals("id", id)
		return ctx.Next()
	}
}
