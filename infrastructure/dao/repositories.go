package dao

import (
	"gorm.io/gorm"

	"go-api-cms/domain/repository"
	"go-api-cms/infrastructure/persistence"
)

type Repositories struct {
	db                 *gorm.DB
	UserRepo           repository.UserRepository
	RoleRepo           repository.RoleRepository
	PermissionRepo     repository.PermissionRepository
	ArticleRepo        repository.ArticleRepository
	ArticleVersionRepo repository.ArticleVersionRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		db:                 db,
		UserRepo:           persistence.NewUserDBRepository(db),
		RoleRepo:           persistence.NewRoleDBRepository(db),
		PermissionRepo:     persistence.NewPermissionDBRepository(db),
		ArticleRepo:        persistence.NewArticleDBRepository(db),
		ArticleVersionRepo: persistence.NewArticleVersionDBRepository(db),
	}
}

func (r *Repositories) Tx() (*Repositories, *gorm.DB) {
	tx := r.db.Begin()
	return NewRepositories(tx), tx
}
