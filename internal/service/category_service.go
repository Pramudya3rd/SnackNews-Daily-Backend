package service

import (
	"news-shared-service/internal/models"
	"news-shared-service/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(cat *models.Category) error {
	return s.repo.Create(cat)
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAll()
}
