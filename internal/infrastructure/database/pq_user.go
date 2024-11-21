package database

import (
	"auth-service/internal/entities"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	DB *gorm.DB
}

func (r *PostgresUserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Save(user *entities.User) error {
	return r.DB.Create(user).Error
}
