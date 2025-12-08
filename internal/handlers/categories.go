package handlers

import (
	"net/http"

	"news-shared-service/internal/models"
	svc "news-shared-service/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
    service *svc.CategoryService
}

func NewCategoryHandler(s *svc.CategoryService) *CategoryHandler {
    return &CategoryHandler{service: s}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
    var payload models.Category
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.CreateCategory(&payload); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
        return
    }

    c.JSON(http.StatusCreated, payload)
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
    categories, err := h.service.GetAllCategories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, categories)
}