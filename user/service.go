package user

type userService struct {
	repository IUserRepository
}

type IUserService interface {
	Create(createUserRequest CreateUserRequest) (*User, error)
}

func NewUserService(repository IUserRepository) IUserService {
	return &userService{repository}
}

func (service *userService) Create(createUserRequest CreateUserRequest) (*User, error) {
	user := &User{
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
	}

	createdUser, err := service.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
