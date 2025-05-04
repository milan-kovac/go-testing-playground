package user_test

import (
	"errors"
	"testing"

	"github.com/milan-kovac/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user.NewUserService(mockRepo)

	t.Run("successfully creates user", func(t *testing.T) {
		req := user.CreateUserRequest{
			FirstName: "Milan",
			LastName:  "Kovač",
		}

		expectedUser := &user.User{
			FirstName: "Milan",
			LastName:  "Kovač",
		}

		mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(expectedUser, nil).Once()

		createdUser, err := service.Create(req)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser.FirstName, createdUser.FirstName)
		assert.Equal(t, expectedUser.LastName, createdUser.LastName)

		mockRepo.AssertExpectations(t)
	})

	t.Run("fails to create user when repo returns error", func(t *testing.T) {
		req := user.CreateUserRequest{
			FirstName: "Error",
			LastName:  "Case",
		}

		mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(nil, errors.New("db error")).Once()

		createdUser, err := service.Create(req)

		assert.Nil(t, createdUser)
		assert.EqualError(t, err, "db error")

		mockRepo.AssertExpectations(t)
	})
}
