package repositories

import "auth-service/internal/entities"

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	Save(user *entities.User) error
}
