package usecases

import (
	"auth-service/internal/entities"
	"auth-service/internal/interfaces/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserUseCase_Execute(t *testing.T) {
	repo := repositories.NewMockUserRepository()
	useCase := &RegisterUserUseCase{UserRepo: repo}

	t.Run("Successful registration", func(t *testing.T) {
		user := &entities.User{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "Test User",
		}

		err := useCase.Execute(user)
		assert.NoError(t, err)
		assert.NotEmpty(t, user.Password)
	})

	t.Run("User already exists", func(t *testing.T) {
		user := &entities.User{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "Test User",
		}

		err := useCase.Execute(user)
		assert.Error(t, err)
		assert.Equal(t, "user already exists", err.Error())
	})

	t.Run("Invalid input", func(t *testing.T) {
		user := &entities.User{
			Email:    "",
			Password: "",
		}

		err := useCase.Execute(user)
		assert.Error(t, err)
		assert.Equal(t, "email and password are required", err.Error())
	})
}
