package handlers

import (
	"net/http"
	"time"

	"news-shared-service/internal/models"
	"news-shared-service/internal/service"

	"github.com/gin-gonic/gin"
)

// NewsHandler struct to hold the news service
type NewsHandler struct {
	service *service.NewsService
}

// NewNewsHandler creates a new instance of NewsHandler
func NewNewsHandler(service *service.NewsService) *NewsHandler {
	return &NewsHandler{service: service}
}

// CreateNews handles the creation of a news item
func (h *NewsHandler) CreateNews(c *gin.Context) {
	// Accept frontend payload which may include numeric timestamps and extra fields.
	var req struct {
		ID             string `json:"id"`
		Title          string `json:"title"`
		Content        string `json:"content"`
		Author         string `json:"author"`
		Category       string `json:"category"`
		Image          string `json:"image"`
		SourceURL      string `json:"sourceUrl"`
		DisplaySection string `json:"displaySection"`
		Archived       bool   `json:"archived"`
		CreatedAt      *int64 `json:"createdAt"`
		UpdatedAt      *int64 `json:"updatedAt"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	news := models.News{
		ID:             req.ID,
		Title:          req.Title,
		Content:        req.Content,
		Author:         req.Author,
		Category:       req.Category,
		Image:          req.Image,
		SourceURL:      req.SourceURL,
		DisplaySection: req.DisplaySection,
		Archived:       req.Archived,
	}

	if req.CreatedAt != nil {
		news.CreatedAt = time.UnixMilli(*req.CreatedAt)
	}
	if req.UpdatedAt != nil {
		news.UpdatedAt = time.UnixMilli(*req.UpdatedAt)
	}

	if err := h.service.CreateNews(&news); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, news)
}

// GetNews retrieves a news item by ID
func (h *NewsHandler) GetNews(c *gin.Context) {
	id := c.Param("id")
	news, err := h.service.GetNews(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, news)
}

// UpdateNews handles the update of a news item
func (h *NewsHandler) UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Title          string `json:"title"`
		Content        string `json:"content"`
		Author         string `json:"author"`
		Category       string `json:"category"`
		Image          string `json:"image"`
		SourceURL      string `json:"sourceUrl"`
		DisplaySection string `json:"displaySection"`
		Archived       bool   `json:"archived"`
		UpdatedAt      *int64 `json:"updatedAt"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fetch existing
	existing, err := h.service.GetNews(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	// apply updates
	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.Content != "" {
		existing.Content = req.Content
	}
	if req.Author != "" {
		existing.Author = req.Author
	}
	if req.Category != "" {
		existing.Category = req.Category
	}
	if req.Image != "" {
		existing.Image = req.Image
	}
	if req.SourceURL != "" {
		existing.SourceURL = req.SourceURL
	}
	if req.DisplaySection != "" {
		existing.DisplaySection = req.DisplaySection
	}
	existing.Archived = req.Archived
	if req.UpdatedAt != nil {
		existing.UpdatedAt = time.UnixMilli(*req.UpdatedAt)
	}

	if err := h.service.UpdateNews(existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existing)
}

// DeleteNews handles the deletion of a news item
func (h *NewsHandler) DeleteNews(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteNews(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetAllNews retrieves all news items
func (h *NewsHandler) GetAllNews(c *gin.Context) {
	newsList, err := h.service.GetAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newsList)
}