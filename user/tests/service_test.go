package user_test

import (
	"errors"
	"testing"
	"time"

	"github.com/milan-kovac/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user.NewUserService(mockRepo)

	t.Run("should create user successfully", func(t *testing.T) {
		req := user.CreateUserRequest{
			FirstName: "James",
			LastName:  "Smith",
		}

		expectedUser := &user.User{
			ID:        1,
			FirstName: "James",
			LastName:  "Smith",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(expectedUser, nil).Once()

		createdUser, err := service.Create(req)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser.FirstName, createdUser.FirstName)
		assert.Equal(t, expectedUser.LastName, createdUser.LastName)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to create user when repo returns error", func(t *testing.T) {
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

	t.Run("should return all users successfully", func(t *testing.T) {
		expectedUsers := []user.User{
			{
				ID:        1,
				FirstName: "Michael",
				LastName:  "Johnson",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        2,
				FirstName: "Emily",
				LastName:  "Davis",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		mockRepo.On("GetAll").Return(expectedUsers, nil).Once()

		users, err := service.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when failing to get all users", func(t *testing.T) {
		mockRepo.On("GetAll").Return([]user.User(nil), errors.New("db error")).Once()

		users, err := service.GetAll()

		assert.Nil(t, users)
		assert.EqualError(t, err, "db error")

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return user successfully", func(t *testing.T) {
		expectedUser := user.User{
			ID:        1,
			FirstName: "Michael",
			LastName:  "Johnson",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mockRepo.On("Get", int(expectedUser.ID)).Return(&expectedUser, nil).Once()

		user, err := service.Get(int(expectedUser.ID))

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, *user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when failing to get user", func(t *testing.T) {
		mockRepo.On("Get", 1234).Return((*user.User)(nil), errors.New("db error")).Once()

		user, err := service.Get(1234)

		assert.Nil(t, user)
		assert.EqualError(t, err, "db error")

		mockRepo.AssertExpectations(t)
	})
}
