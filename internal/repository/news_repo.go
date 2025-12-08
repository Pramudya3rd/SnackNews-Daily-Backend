package repository

import (
	"news-shared-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(news *models.News) error
	GetByID(id string) (*models.News, error)
	GetAll() ([]models.News, error)
	Update(news *models.News) error
	Delete(id string) error
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) Create(news *models.News) error {
	if news.ID == "" {
		news.ID = uuid.NewString()
	}
	return r.db.Create(news).Error
}

func (r *newsRepository) GetByID(id string) (*models.News, error) {
	var news models.News
	if err := r.db.First(&news, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) GetAll() ([]models.News, error) {
	var news []models.News
	if err := r.db.Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) Update(news *models.News) error {
	return r.db.Save(news).Error
}

func (r *newsRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.News{}).Error
}