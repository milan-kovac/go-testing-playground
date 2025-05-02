package user

type CreateUserRequest struct {
	FirstName string `json:"firstName" validate:"required,min=3"`
	LastName  string `json:"lastName" validate:"required,min=3"`
}
