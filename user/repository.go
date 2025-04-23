package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) create() {

}
