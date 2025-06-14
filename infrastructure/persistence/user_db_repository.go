package persistence

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go-api-cms/domain/entity"
)

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (u *UserDBRepository) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := u.db.WithContext(ctx).First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDBRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := u.db.WithContext(ctx).First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDBRepository) Create(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *UserDBRepository) Update(ctx context.Context, userID uuid.UUID, user *entity.User) error {
	q := u.db.WithContext(ctx).Where("id = ?", userID).Updates(user)
	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return q.Error
}

func (u *UserDBRepository) AssignRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error {
	var userRoles []entity.UserRole

	for _, roleID := range roleIDs {
		userRoles = append(userRoles, entity.UserRole{
			UserID: userID,
			RoleID: roleID,
		})
	}

	if len(userRoles) == 0 {
		return nil
	}

	return u.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "role_id"}},
		DoNothing: true,
	}).Create(&userRoles).Error
}

func (u *UserDBRepository) AssignPermissions(ctx context.Context, userID uuid.UUID, permissionIDs []uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
