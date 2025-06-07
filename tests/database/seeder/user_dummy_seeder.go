package seeder

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"go-api-cms/domain/entity"
)

type UserDummySeeder struct {
	Users []*entity.User
}

func (u *UserDummySeeder) Seed(ctx context.Context, db *gorm.DB) error {
	users := []*entity.User{
		{
			ID:       uuid.New(),
			Name:     "John Doe",
			Username: "john.doe",
			Email:    "john.doe@mailinator.com",
			Password: "secret",
			IsActive: true,
		},
	}

	if err := db.WithContext(ctx).Create(&users).Error; err != nil {
		return err
	}

	u.Users = users

	return nil
}
