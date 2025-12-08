package repository

import (
	"news-shared-service/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
    Create(category *models.Category) error
    GetAll() ([]models.Category, error)
}

type categoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
    return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *models.Category) error {
    return r.db.Create(category).Error
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
    var categories []models.Category
    if err := r.db.Find(&categories).Error; err != nil {
        return nil, err
    }
    return categories, nil
}
