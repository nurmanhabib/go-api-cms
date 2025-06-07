package repository

import (
	"context"

	"github.com/google/uuid"

	"go-api-cms/domain/entity"
)

type RoleRepository interface {
	FindAll(ctx context.Context) ([]*entity.Role, error)
	FindById(ctx context.Context, id int64) (*entity.Role, error)
	FindByNames(ctx context.Context, names []string) ([]*entity.Role, error)
	Create(ctx context.Context, role *entity.Role) error
	Update(ctx context.Context, roleID uuid.UUID, role *entity.Role) error
	AssignPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error
}
