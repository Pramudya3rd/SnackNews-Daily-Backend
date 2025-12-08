package main

import (
	"fmt"
	"news-shared-service/internal/api"
	"news-shared-service/internal/config"
	"news-shared-service/internal/models"
	"news-shared-service/internal/repository"
	"news-shared-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        panic(err)
    }

    // Initialize database connection (GORM)
    gormDB, err := repository.NewDBFromConfig(cfg.Database)
    if err != nil {
        panic(err)
    }

    // Auto-migrate models (creates tables if they do not exist)
    if err := gormDB.AutoMigrate(&models.News{}, &models.Category{}, &models.User{}); err != nil {
        panic(err)
    }

    // Initialize repositories and services
    newsRepo := repository.NewNewsRepository(gormDB)
    newsService := service.NewNewsService(newsRepo)
    categoryRepo := repository.NewCategoryRepository(gormDB)
    categoryService := service.NewCategoryService(categoryRepo)

    // Initialize auth repo/service and ensure default admin user exists
    userRepo := repository.NewUserRepository(gormDB)
    authService := service.NewAuthService(userRepo, cfg.Server.JWTSecret)

    // Ensure default admin user (only if not exists). Password: 'admin123' (change in production)
    if _, err := userRepo.GetByUsername("admin"); err != nil {
        // create default admin
        if err := authService.CreateUser("admin", "admin123"); err != nil {
            panic(err)
        }
    }

    // Set up Gin engine
    r := gin.Default()

    // Set up routes
    api.SetupRoutes(r, newsService, categoryService, authService)

    // Start the server
    if err := r.Run(":" + fmt.Sprintf("%d", cfg.Server.Port)); err != nil {
        panic(err)
    }
}