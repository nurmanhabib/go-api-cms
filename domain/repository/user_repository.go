package repository

import (
	"context"

	"github.com/google/uuid"

	"go-api-cms/domain/entity"
)

type UserRepository interface {
	FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, userID uuid.UUID, user *entity.User) error
	AssignRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error
	AssignPermissions(ctx context.Context, userID uuid.UUID, permissionIDs []uuid.UUID) error
}
