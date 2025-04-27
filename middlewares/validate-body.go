package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/milan-kovac/common"
)

var validate = validator.New()

func ValidateBody[T any]() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body T

		if err := ctx.BodyParser(&body); err != nil {
			return common.ErrorResponse(ctx, fiber.ErrBadRequest.Code, err)
		}

		if err := validate.Struct(body); err != nil {
			return common.ErrorResponse(ctx, fiber.ErrBadRequest.Code, err)
		}

		ctx.Locals("body", body)

		return ctx.Next()
	}
}
