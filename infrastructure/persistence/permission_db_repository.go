package persistence

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"go-api-cms/domain/entity"
)

type PermissionDBRepository struct {
	db *gorm.DB
}

func NewPermissionDBRepository(db *gorm.DB) *PermissionDBRepository {
	return &PermissionDBRepository{db: db}
}

func (p *PermissionDBRepository) FindAllByRoleID(ctx context.Context, roleID uuid.UUID) ([]entity.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PermissionDBRepository) FindAllByPermissionGroupID(ctx context.Context, permissionGroupID uuid.UUID) ([]entity.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PermissionDBRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error) {
	//TODO implement me
	panic("implement me")
}
