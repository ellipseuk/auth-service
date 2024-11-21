package usecases

import (
	"auth-service/internal/entities"
	"auth-service/internal/interfaces/repositories"
	"auth-service/pkg/hash"
	"errors"
	"fmt"
)

type RegisterUserUseCase struct {
	UserRepo repositories.UserRepository
}

func (uc *RegisterUserUseCase) Execute(user *entities.User) error {
	fmt.Printf("Registering user: %v\n", user)

	// Check if email and password are provided
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	// Check if user already exists
	existingUser, _ := uc.UserRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	// Save user
	return uc.UserRepo.Save(user)
}
