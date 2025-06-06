package repository

import (
	"context"

	"go-api-cms/domain/entity"
)

type AuthRepository interface {
	GenerateToken(user *entity.User) (string, error)
	ParseToken(ctx context.Context, tokenString string) (*entity.User, error)
}
