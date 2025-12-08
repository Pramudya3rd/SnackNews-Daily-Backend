package service

import (
	"errors"
	"news-shared-service/internal/models"
	"news-shared-service/internal/repository"
)

type NewsService struct {
	repo repository.NewsRepository
}

func NewNewsService(repo repository.NewsRepository) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) CreateNews(news *models.News) error {
	if news == nil {
		return errors.New("news cannot be nil")
	}
	return s.repo.Create(news)
}

func (s *NewsService) GetNews(id string) (*models.News, error) {
	return s.repo.GetByID(id)
}

func (s *NewsService) UpdateNews(news *models.News) error {
	if news == nil {
		return errors.New("news cannot be nil")
	}
	return s.repo.Update(news)
}

func (s *NewsService) DeleteNews(id string) error {
	return s.repo.Delete(id)
}

func (s *NewsService) GetAllNews() ([]models.News, error) {
	return s.repo.GetAll()
}