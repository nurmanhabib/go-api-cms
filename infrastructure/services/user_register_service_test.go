package services_test

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-api-cms/domain/entity"
	"go-api-cms/infrastructure/services"
	"go-api-cms/tests"
)

func TestUserRegisterService_Registration(t *testing.T) {
	suite := tests.NewSuite()
	ctx := suite.Ctx

	t.Run("registration as editor", func(t *testing.T) {
		newUser := &entity.User{
			ID:       uuid.New(),
			Name:     faker.Name(),
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: faker.Password(),
			IsActive: true,
		}

		assignedRoles := []string{"editor"}

		svc := services.NewUserRegisterService(suite.App)
		err := svc.Registration(ctx, newUser, assignedRoles)
		require.NoError(t, err)
		assert.NotEmpty(t, newUser.ID)

		t.Run("check from database", func(t *testing.T) {
			var checkUser entity.User
			errCheck := suite.App.DB.WithContext(ctx).First(&checkUser, "id = ?", newUser.ID).Error
			require.NoError(t, errCheck)
			assert.Equal(t, newUser.Name, checkUser.Name)
			assert.Equal(t, newUser.Email, checkUser.Email)
		})

		t.Run("check assigned roles", func(t *testing.T) {
			var checkRoles []*entity.Role

			errCheck := suite.App.DB.WithContext(ctx).
				Joins("JOIN user_roles ON user_roles.role_id = roles.id").
				Where("user_roles.user_id = ?", newUser.ID).
				Find(&checkRoles).
				Error

			require.NoError(t, errCheck)
			require.Len(t, checkRoles, 1)
			assert.Equal(t, "editor", checkRoles[0].Name)
		})
	})

	t.Run("registration as unknown roles", func(t *testing.T) {
		newUser := &entity.User{
			ID:       uuid.New(),
			Name:     faker.Name(),
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: faker.Password(),
			IsActive: true,
		}

		assignedRoles := []string{"unknown_roles"}

		svc := services.NewUserRegisterService(suite.App)
		err := svc.Registration(ctx, newUser, assignedRoles)
		require.NoError(t, err)
		assert.NotEmpty(t, newUser.ID)

		t.Run("check from database", func(t *testing.T) {
			var checkUser entity.User
			errCheck := suite.App.DB.WithContext(ctx).First(&checkUser, "id = ?", newUser.ID).Error
			require.NoError(t, errCheck)
			assert.Equal(t, newUser.Name, checkUser.Name)
			assert.Equal(t, newUser.Username, checkUser.Username)
			assert.Equal(t, newUser.Email, checkUser.Email)
		})

		t.Run("check assigned roles", func(t *testing.T) {
			var checkRoles []*entity.Role

			errCheck := suite.App.DB.WithContext(ctx).
				Joins("JOIN user_roles ON user_roles.role_id = roles.id").
				Where("user_roles.user_id = ?", newUser.ID).
				Find(&checkRoles).
				Error

			require.NoError(t, errCheck)
			require.Len(t, checkRoles, 0)
		})
	})
}
