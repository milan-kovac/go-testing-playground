package user

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Create(user *User) (*User, error)
	GetAll() ([]User, error)
	Get(id int) (*User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (repository *userRepository) Create(user *User) (*User, error) {
	if err := repository.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *userRepository) GetAll() ([]User, error) {
	var users []User

	if err := repository.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *userRepository) Get(id int) (*User, error) {
	var user User

	if err := repository.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
