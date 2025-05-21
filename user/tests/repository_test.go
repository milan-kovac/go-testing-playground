package user_test

import (
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/milan-kovac/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:test.db?_pragma=foreign_keys(1)"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the test database")
	}
	db.AutoMigrate(&user.User{})
	return db
}

func TestUserRepository(t *testing.T) {
	db := SetupTestDB()
	repo := user.NewUserRepository(db)

	t.Run("should create user successfully", func(t *testing.T) {
		user := &user.User{
			FirstName: "James",
			LastName:  "Smith",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		createdUser, err := repo.Create(user)

		assert.NoError(t, err)
		assert.NotZero(t, createdUser.ID)
		assert.Equal(t, "James", createdUser.FirstName)
	})

	t.Run("should get all users successfully", func(t *testing.T) {
		usersToCreate := []user.User{
			{FirstName: "Alice", LastName: "Wonderland"},
			{FirstName: "Bob", LastName: "Builder"},
		}

		for i := range usersToCreate {
			_, err := repo.Create(&usersToCreate[i])
			assert.NoError(t, err)
		}

		users, err := repo.GetAll()

		assert.NoError(t, err)

		found := make(map[int]bool, len(usersToCreate))

		for i, expectedUser := range usersToCreate {
			for _, actualUser := range users {
				if actualUser.FirstName == expectedUser.FirstName && actualUser.LastName == expectedUser.LastName {
					found[i] = true
					break
				}
			}
		}

		for i := range usersToCreate {
			assert.True(t, found[i], "Expected user not found: %+v", usersToCreate[i])
		}
	})

	t.Run("should get user by ID successfully", func(t *testing.T) {
		user := &user.User{
			FirstName: "Charlie",
			LastName:  "Brown",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		createdUser, err := repo.Create(user)
		assert.NoError(t, err)
		assert.NotZero(t, createdUser.ID)

		fetchedUser, err := repo.Get(int(createdUser.ID))
		assert.NoError(t, err)
		assert.NotNil(t, fetchedUser)
		assert.Equal(t, createdUser.ID, fetchedUser.ID)
		assert.Equal(t, user.FirstName, fetchedUser.FirstName)
		assert.Equal(t, user.LastName, fetchedUser.LastName)
	})

	t.Run("should return error when user not found", func(t *testing.T) {
		fetchedUser, err := repo.Get(99999)
		assert.Error(t, err)
		assert.Nil(t, fetchedUser)
	})

}
