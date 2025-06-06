package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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

	svc := services.NewAuthService(suite.App)
	token, err := svc.GenerateToken(userDummy.Users[0])
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
