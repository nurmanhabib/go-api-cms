package dao

import (
	"gorm.io/gorm"

	"go-api-cms/domain/repository"
	"go-api-cms/infrastructure/persistence"
)

type Repositories struct {
	UserRepo       repository.UserRepository
	RoleRepo       repository.RoleRepository
	PermissionRepo repository.PermissionRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo:       persistence.NewUserDBRepository(db),
		RoleRepo:       persistence.NewRoleDBRepository(db),
		PermissionRepo: persistence.NewPermissionDBRepository(db),
	}
}
