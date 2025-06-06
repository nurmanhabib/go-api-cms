package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-api-cms/infrastructure/exception"
	"go-api-cms/infrastructure/services"
	"go-api-cms/tests"
	"go-api-cms/tests/database/seeder"
)

func TestUserAuthService_Login(t *testing.T) {
	suite := tests.NewSuite()
	ctx := suite.Ctx
	userDummy := &seeder.UserDummySeeder{}

	if errSeed := suite.RunSeeder(userDummy); errSeed != nil {
		t.Error(errSeed)
	}

	t.Run("given valid credentials", func(t *testing.T) {
		svc := services.NewUserAuthService(suite.App)
		loggedInUser, err := svc.Login(ctx, "john.doe", "secret")

		require.NoError(t, err)
		require.NotNil(t, loggedInUser)
		assert.Equal(t, loggedInUser.Username, "john.doe")
	})

	t.Run("given invalid credentials", func(t *testing.T) {
		svc := services.NewUserAuthService(suite.App)
		loggedInUser, err := svc.Login(ctx, "john.doe", "wrong_password")

		require.Error(t, err)
		require.Nil(t, loggedInUser)

		var svcException *exception.ServiceError

		if assert.ErrorAs(t, err, &svcException) {
			require.Len(t, svcException.Violations, 1)
			assert.Equal(t, svcException.Violations[0].Field, "username")
		}
	})
}
