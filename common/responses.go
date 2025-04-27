package common

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ErrorDetails struct {
	Code    int    `json:"code"`
	Error   string `json:"error,omitempty"`
	TraceID string `json:"traceId"`
	Method  string `json:"method"`
}

func ErrorResponse(ctx *fiber.Ctx, code int, err error) error {
	traceID := uuid.New().String()
	method := ctx.Method()

	return ctx.Status(code).JSON(ErrorDetails{
		Code:    code,
		Error:   formatError(err),
		TraceID: traceID,
		Method:  method,
	})
}

func formatError(err error) string {
	var errorMessages []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' validation failed", e.Field()))
		}
	} else {
		errorMessages = append(errorMessages, err.Error())
	}

	return strings.Join(errorMessages, ", ")
}
