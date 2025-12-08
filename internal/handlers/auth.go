package handlers

import (
	"log"
	"net/http"

	"news-shared-service/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
    svc *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
    return &AuthHandler{svc: s}
}

type loginRequest struct{
    Username string `json:"username"`
    Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req loginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Login Error: invalid JSON request - %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
        return
    }

    // Validate input
    if req.Username == "" || req.Password == "" {
        log.Printf("Login Error: empty username or password")
        c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
        return
    }

    log.Printf("Login Attempt: user %s", req.Username)
    token, err := h.svc.Authenticate(req.Username, req.Password)
    if err != nil {
        log.Printf("Login Failed: user %s - %v", req.Username, err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
        return
    }

    log.Printf("Login Success: user %s", req.Username)
    c.JSON(http.StatusOK, gin.H{"token": token})
}
