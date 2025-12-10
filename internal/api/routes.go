package api

import (
	"news-shared-service/internal/handlers"
	"news-shared-service/internal/service"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRoutes(router *gin.Engine, newsService *service.NewsService, categoryService *service.CategoryService, authService *service.AuthService) {
    // Apply CORS middleware
    router.Use(CORSMiddleware())
    
    // Health check endpoint (no prefix)
    router.GET("/health", handlers.HealthCheck)

    newsHandler := handlers.NewNewsHandler(newsService)
    uploadHandler := handlers.NewUploadHandler()

    // API v1 routes with /api prefix
    apiGroup := router.Group("/api")
    {
        // News endpoints
        newsGroup := apiGroup.Group("/news")
        {
            // Public read endpoints
            newsGroup.GET("/", newsHandler.GetAllNews)
            newsGroup.GET("/:id", newsHandler.GetNews)

            // Protected write endpoints
            newsGroup.POST("/", AuthMiddleware(authService), newsHandler.CreateNews)
            newsGroup.PUT("/:id", AuthMiddleware(authService), newsHandler.UpdateNews)
            newsGroup.DELETE("/:id", AuthMiddleware(authService), newsHandler.DeleteNews)
        }

        // Categories endpoints
        categoryHandler := handlers.NewCategoryHandler(categoryService)
        categoriesGroup := apiGroup.Group("/categories")
        {
            categoriesGroup.GET("/", categoryHandler.GetCategories)
            // protect category creation
            categoriesGroup.POST("/", AuthMiddleware(authService), categoryHandler.CreateCategory)
        }

        // Upload endpoints
        uploadsGroup := apiGroup.Group("/uploads")
        {
            // Image upload endpoint (protected)
            uploadsGroup.POST("/images", AuthMiddleware(authService), uploadHandler.UploadImage)
            // Serve images (public)
            uploadsGroup.GET("/images/:filename", uploadHandler.ServeImage)
            // Delete image endpoint (protected)
            uploadsGroup.DELETE("/images/:filename", AuthMiddleware(authService), uploadHandler.DeleteImage)
        }

        // Auth endpoints
        authHandler := handlers.NewAuthHandler(authService)
        authGroup := apiGroup.Group("/auth")
        {
            authGroup.POST("/login", authHandler.Login)
        }
    }
}