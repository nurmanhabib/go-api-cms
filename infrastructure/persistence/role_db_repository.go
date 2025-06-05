package persistence

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"go-api-cms/domain/entity"
)

type RoleDBRepository struct {
	db *gorm.DB
}

func NewRoleDBRepository(db *gorm.DB) *RoleDBRepository {
	return &RoleDBRepository{db: db}
}

func (r *RoleDBRepository) FindAll(ctx context.Context) ([]*entity.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RoleDBRepository) FindById(ctx context.Context, id int64) (*entity.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RoleDBRepository) Create(ctx context.Context, role *entity.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r *RoleDBRepository) Update(ctx context.Context, roleID uuid.UUID, role *entity.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r *RoleDBRepository) AssignPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
