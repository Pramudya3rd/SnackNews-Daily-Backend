package service

import (
	"errors"
	"time"

	"news-shared-service/internal/models"
	"news-shared-service/internal/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    repo   repository.UserRepository
    secret string
}

func NewAuthService(repo repository.UserRepository, secret string) *AuthService {
    return &AuthService{repo: repo, secret: secret}
}

func (s *AuthService) Authenticate(username, password string) (string, error) {
    user, err := s.repo.GetByUsername(username)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    // create JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "usr": user.Username,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
    })

    signed, err := token.SignedString([]byte(s.secret))
    if err != nil {
        return "", err
    }
    return signed, nil
}

func (s *AuthService) CreateUser(username, password string) error {
    // hash password
    h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u := &models.User{Username: username, Password: string(h)}
    return s.repo.Create(u)
}

// ValidateToken parses and validates a JWT token and returns the user id and username
func (s *AuthService) ValidateToken(tokenStr string) (uint, string, error) {
    if tokenStr == "" {
        return 0, "", errors.New("empty token")
    }
    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(s.secret), nil
    })
    if err != nil || !token.Valid {
        return 0, "", errors.New("invalid token")
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        // sub stored as numeric or string; handle both
        var uid uint
        switch v := claims["sub"].(type) {
        case float64:
            uid = uint(v)
        case int:
            uid = uint(v)
        case string:
            // try parse
            // ignore parse errors and leave uid 0
        }
        usr, _ := claims["usr"].(string)
        return uid, usr, nil
    }
    return 0, "", errors.New("invalid claims")
}
