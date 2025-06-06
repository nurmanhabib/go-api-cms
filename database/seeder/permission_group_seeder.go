package seeder

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go-api-cms/domain/entity"
)

const (
	PermissionGroupUserManagement    = "ca5ce3c9-9f33-4e60-aca3-454ddac2d9be"
	PermissionGroupArticleManagement = "4df9b8e9-8dbb-4fc3-916c-5e394dba12ae"
)

type PermissionGroupSeeder struct{}

func NewPermissionGroupSeeder() *PermissionGroupSeeder {
	return &PermissionGroupSeeder{}
}

func (p *PermissionGroupSeeder) Seed(ctx context.Context, db *gorm.DB) error {
	permissionGroups := []*entity.PermissionGroup{
		{
			ID:    uuid.MustParse(PermissionGroupUserManagement),
			Name:  "user_management",
			Label: "User Management",
		},
		{
			ID:    uuid.MustParse(PermissionGroupArticleManagement),
			Name:  "article_management",
			Label: "Article Management",
		},
	}

	for _, permissionGroup := range permissionGroups {
		if err := db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(permissionGroup).Error; err != nil {
			return err
		}
	}

	return nil
}
