package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/milan-kovac/common"
)

type IUserController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
}

type userController struct {
	userService IUserService
}

func NewUserController(userService IUserService) IUserController {
	return &userController{userService}
}

func (controller *userController) Create(ctx *fiber.Ctx) error {
	var createUserRequest CreateUserRequest = ctx.Locals("body").(CreateUserRequest)

	createdUser, err := controller.userService.Create(createUserRequest)

	if err != nil {
		return common.ErrorResponse(ctx, fiber.ErrConflict.Code, err)
	}

	return common.SuccessResponse(ctx, fiber.StatusCreated, "User created.", createdUser)

}

func (controller *userController) GetAll(ctx *fiber.Ctx) error {
	users, err := controller.userService.GetAll()

	if err != nil {
		return common.ErrorResponse(ctx, fiber.ErrConflict.Code, err)
	}

	return common.SuccessResponse(ctx, fiber.StatusCreated, "Users fetched successfully.", users)
}

func (controller *userController) Get(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	user, err := controller.userService.Get(id)

	if err != nil {
		return common.ErrorResponse(ctx, fiber.ErrConflict.Code, err)
	}

	return common.SuccessResponse(ctx, fiber.StatusCreated, "User fetched successfully.", user)
}
