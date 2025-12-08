package api

import (
	"log"
	"net/http"

	"news-shared-service/internal/service"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs the details of each request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL)
		c.Next()
		log.Printf("Response: %d", c.Writer.Status())
	}
}

// AuthMiddleware returns a middleware that validates Bearer JWT using provided AuthService
func AuthMiddleware(authSvc *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			log.Printf("Auth Error: missing authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}
		// expect "Bearer <token>"
		var token string
		if len(header) > 7 && header[:7] == "Bearer " {
			token = header[7:]
		} else {
			token = header
		}

		uid, username, err := authSvc.ValidateToken(token)
		if err != nil {
			log.Printf("Auth Error: invalid token - %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		// set values into context for handlers
		c.Set("userID", uid)
		c.Set("username", username)
		log.Printf("Auth Success: user %s (ID: %d)", username, uid)
		c.Next()
	}
}