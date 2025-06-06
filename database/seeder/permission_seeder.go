package seeder

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go-api-cms/domain/entity"
)

const (
	PermissionUserCreate     = "ca5ce3c9-9f33-4e60-aca3-454ddac2d9be"
	PermissionUserEdit       = "4df9b8e9-8dbb-4fc3-916c-5e394dba12ae"
	PermissionUserDelete     = "a90dd935-6cb5-4461-86e3-6533887ab796"
	PermissionUserAssignRole = "1f61c9a3-3148-462f-813d-02f43cf8e1df"

	PermissionArticleCreate  = "e062ce95-e7c9-43cd-8c90-14f550b22a82"
	PermissionArticleEdit    = "6f2e4a77-5d5e-4394-b7d0-ec1c1caa7f48"
	PermissionArticleDelete  = "9febfd89-cf8e-4fa6-ab46-4f735d4e3a8d"
	PermissionArticlePublish = "07829d21-fc65-4e0d-97e9-9f14b4c44c3f"
)

type PermissionSeeder struct{}

func NewPermissionSeeder() *PermissionSeeder {
	return &PermissionSeeder{}
}

func (p *PermissionSeeder) Seed(ctx context.Context, db *gorm.DB) error {
	permissions := []*entity.Permission{
		{
			ID:      uuid.MustParse(PermissionUserCreate),
			GroupID: uuid.MustParse(PermissionGroupUserManagement),
			Name:    "user.create",
			Label:   "Create User",
		},
		{
			ID:      uuid.MustParse(PermissionUserEdit),
			GroupID: uuid.MustParse(PermissionGroupUserManagement),
			Name:    "user.edit",
			Label:   "Edit User",
		},
		{
			ID:      uuid.MustParse(PermissionUserDelete),
			GroupID: uuid.MustParse(PermissionGroupUserManagement),
			Name:    "user.delete",
			Label:   "Delete User",
		},
		{
			ID:      uuid.MustParse(PermissionUserAssignRole),
			GroupID: uuid.MustParse(PermissionGroupUserManagement),
			Name:    "user.assign_role",
			Label:   "Assign Role to User",
		},
		{
			ID:      uuid.MustParse(PermissionArticleCreate),
			GroupID: uuid.MustParse(PermissionGroupArticleManagement),
			Name:    "article.create",
			Label:   "Create Article",
		},
		{
			ID:      uuid.MustParse(PermissionArticleEdit),
			GroupID: uuid.MustParse(PermissionGroupArticleManagement),
			Name:    "article.edit",
			Label:   "Edit Article",
		},
		{
			ID:      uuid.MustParse(PermissionArticleDelete),
			GroupID: uuid.MustParse(PermissionGroupArticleManagement),
			Name:    "article.delete",
			Label:   "Delete Article",
		},
		{
			ID:      uuid.MustParse(PermissionArticlePublish),
			GroupID: uuid.MustParse(PermissionGroupArticleManagement),
			Name:    "article.publish",
			Label:   "Publish Article",
		},
	}

	for _, permission := range permissions {
		if err := db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(permission).Error; err != nil {
			return err
		}
	}

	return nil
}
