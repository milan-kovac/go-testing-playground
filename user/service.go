package user

type userService struct {
	repository IUserRepository
}
type IUserService interface {
	Create(createUserRequest CreateUserRequest) (*User, error)
	GetAll() ([]User, error)
	Get(id int) (*User, error)
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

func (service *userService) GetAll() ([]User, error) {
	users, err := service.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *userService) Get(id int) (*User, error) {
	user, err := service.repository.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
