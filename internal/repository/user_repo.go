package repository

import (
	"news-shared-service/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
    GetByUsername(username string) (*models.User, error)
    Create(user *models.User) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
    var u models.User
    if err := r.db.Where("username = ?", username).First(&u).Error; err != nil {
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}
