package repositories

import (
	"auth-service/internal/entities"
	"errors"
)

type MockUserRepository struct {
	Users map[string]*entities.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{Users: make(map[string]*entities.User)}
}

func (r *MockUserRepository) FindByEmail(email string) (*entities.User, error) {
	if user, exists := r.Users[email]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (r *MockUserRepository) Save(user *entities.User) error {
	if _, exists := r.Users[user.Email]; exists {
		return errors.New("user already exists")
	}
	r.Users[user.Email] = user
	return nil
}
