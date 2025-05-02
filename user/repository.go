package user

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}
type IUserRepository interface {
	Create(user *User) (*User, error)
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
