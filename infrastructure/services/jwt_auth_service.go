package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"go-api-cms/app"
	"go-api-cms/domain/entity"
)

type JwtAuthService struct {
	app       *app.App
	secretKey string
	ttl       time.Duration
}

func NewAuthService(app *app.App) *JwtAuthService {
	return &JwtAuthService{
		app:       app,
		secretKey: app.Config.Auth.JWTSecret,
		ttl:       app.Config.Auth.TTL,
	}
}

func (s *JwtAuthService) GenerateToken(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     time.Now().Add(s.ttl).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JwtAuthService) ParseToken(ctx context.Context, tokenString string) (*entity.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err := s.app.Repo.UserRepo.FindByID(ctx, uuid.MustParse(claims["user_id"].(string)))
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	return nil, errors.New("invalid token")
}
