package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadHandler handles file uploads
type UploadHandler struct{}

// NewUploadHandler creates a new instance of UploadHandler
func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

// UploadImage handles image file uploads
func (h *UploadHandler) UploadImage(c *gin.Context) {
	// Get file from request
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file provided"})
		return
	}

	// Validate file size (max 5MB)
	const maxFileSize = 5 * 1024 * 1024 // 5MB
	if file.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 5MB limit"})
		return
	}

	// Validate file type
	fileExt := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowedExts[fileExt] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only jpg, jpeg, png, gif, and webp files are allowed"})
		return
	}

	// Create uploads directory if it doesn't exist
	uploadsDir := "uploads/images"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Generate unique filename
	uniqueID := uuid.New().String()
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), uniqueID, fileExt)
	filePath := filepath.Join(uploadsDir, filename)

	// Save file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Return the file path that can be used as image URL
	imageURL := fmt.Sprintf("/api/uploads/images/%s", filename)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"url":     imageURL,
		"file":    filename,
	})
}

// ServeImage serves uploaded images
func (h *UploadHandler) ServeImage(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join("uploads/images", filename)

	// Validate path to prevent directory traversal
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file path"})
		return
	}

	absUploadDir, _ := filepath.Abs("uploads/images")
	if len(absPath) < len(absUploadDir) || absPath[:len(absUploadDir)] != absUploadDir {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Check if file exists
	if _, err := os.Stat(filePath); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// Serve the file
	c.File(filePath)
}

// DeleteImage deletes an uploaded image
func (h *UploadHandler) DeleteImage(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join("uploads/images", filename)

	// Validate path to prevent directory traversal
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file path"})
		return
	}

	absUploadDir, _ := filepath.Abs("uploads/images")
	if len(absPath) < len(absUploadDir) || absPath[:len(absUploadDir)] != absUploadDir {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Image deleted successfully",
	})
}
