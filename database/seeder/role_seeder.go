package seeder

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go-api-cms/domain/entity"
)

type RoleSeeder struct{}

func NewRoleSeeder() *RoleSeeder {
	return &RoleSeeder{}
}

type RoleData struct {
	Role        *entity.Role
	Permissions []uuid.UUID
}

const (
	RoleAdmin  = "9fdd9fa9-51d8-4fc2-a04c-d3104e825cdd"
	RoleEditor = "6f31cc0d-3ee8-40f6-bc54-b945be207e72"
)

func (r *RoleSeeder) Seed(ctx context.Context, db *gorm.DB) error {
	roleData := []RoleData{
		{
			Role: &entity.Role{
				ID:   uuid.MustParse(RoleAdmin),
				Name: "admin",
			},
			Permissions: []uuid.UUID{
				uuid.MustParse(PermissionUserCreate),
				uuid.MustParse(PermissionUserEdit),
				uuid.MustParse(PermissionUserDelete),
				uuid.MustParse(PermissionUserAssignRole),
				uuid.MustParse(PermissionArticleCreate),
				uuid.MustParse(PermissionArticleEdit),
				uuid.MustParse(PermissionArticleDelete),
				uuid.MustParse(PermissionArticlePublish),
			},
		},
		{
			Role: &entity.Role{
				ID:   uuid.MustParse(RoleEditor),
				Name: "editor",
			},
			Permissions: []uuid.UUID{
				uuid.MustParse(PermissionArticleCreate),
				uuid.MustParse(PermissionArticleEdit),
				uuid.MustParse(PermissionArticleDelete),
				uuid.MustParse(PermissionArticlePublish),
			},
		},
	}

	for _, role := range roleData {
		q := db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(&role.Role)
		if err := q.Error; err != nil {
			return err
		}

		for _, permissionID := range role.Permissions {
			q2 := db.WithContext(ctx).Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "role_id"}, {Name: "permission_id"}},
				DoNothing: true,
			}).Create(&entity.RolePermission{
				RoleID:       role.Role.ID,
				PermissionID: permissionID,
			})
			if err := q2.Error; err != nil {
				return err
			}
		}
	}

	return nil
}
