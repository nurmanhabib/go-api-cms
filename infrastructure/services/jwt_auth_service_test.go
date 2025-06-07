package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-api-cms/infrastructure/services"
	"go-api-cms/tests"
	"go-api-cms/tests/database/seeder"
)

func TestUserAuthService_GenerateToken(t *testing.T) {
	suite := tests.NewSuite()
	userDummy := &seeder.UserDummySeeder{}

	if errSeed := suite.RunSeeder(userDummy); errSeed != nil {
		t.Error(errSeed)
	}

	t.Run("given user to generate JWT Token", func(t *testing.T) {
		svc := services.NewAuthService(suite.App)
		token, err := svc.GenerateToken(userDummy.Users[0])
		require.NoError(t, err)
		require.NotEmpty(t, token)

		t.Run("while parse token", func(t *testing.T) {
			user, errParse := svc.ParseToken(suite.Ctx, token)
			require.NoError(t, errParse)
			assert.Equal(t, userDummy.Users[0], user)
		})
	})
}
