package repository

import (
	"context"

	"github.com/google/uuid"

	"go-api-cms/domain/entity"
)

type PermissionRepository interface {
	FindAllByRoleID(ctx context.Context, roleID uuid.UUID) ([]entity.Permission, error)
	FindAllByPermissionGroupID(ctx context.Context, permissionGroupID uuid.UUID) ([]entity.Permission, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error)
}
