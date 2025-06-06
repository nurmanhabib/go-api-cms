package services_test

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-api-cms/infrastructure/services"
	"go-api-cms/tests"
	"go-api-cms/tests/database/seeder"
)

func TestArticleService_Create(t *testing.T) {
	suite := tests.NewSuite()
	ctx := suite.Ctx
	userDummy := &seeder.UserDummySeeder{}

	if errSeed := suite.RunSeeder(userDummy); errSeed != nil {
		t.Error(errSeed)
	}

	svc := services.NewArticleService(suite.App)
	article, articleVersion, errCreate := svc.Create(ctx, &services.ArticleRequest{
		Slug:      "my-article-slug",
		Title:     faker.Sentence(),
		Content:   "<p>article content</p>",
		Status:    "draft",
		CreatedBy: userDummy.Users[0].ID,
	})
	require.NoError(t, errCreate)
	require.NotNil(t, article)
	require.NotNil(t, articleVersion)

	assert.Equal(t, "my-article-slug", article.Slug)
	assert.Equal(t, int64(1), articleVersion.Version)

	t.Run("edit article or increment version", func(t *testing.T) {
		editedArticle, editedArticleVersion, errEdit := svc.Update(ctx, article.ID.String(), &services.ArticleRequest{
			Slug:      "new-slug-article",
			Title:     "Edit Title",
			Content:   "<h1>Edit Content</h1>",
			Status:    "published",
			CreatedBy: userDummy.Users[0].ID,
		})

		require.NoError(t, errEdit)
		require.NotNil(t, editedArticle)
		require.NotNil(t, editedArticleVersion)

		assert.Equal(t, int64(2), editedArticleVersion.Version)
		assert.Equal(t, "new-slug-article", editedArticle.Slug)
		assert.Equal(t, "Edit Title", editedArticleVersion.Title)

		t.Run("rollback version test", func(t *testing.T) {
			rollbackArticle, rollbackArticleVersion, errRollback := svc.RollbackVersion(ctx, editedArticle.ID.String(), 1)
			require.NoError(t, errRollback)
			require.NotNil(t, rollbackArticle)
			require.NotNil(t, rollbackArticleVersion)

			assert.Equal(t, articleVersion.ID, rollbackArticleVersion.ID)
			assert.NotEqual(t, "Edit Title", rollbackArticleVersion.Title)
		})
	})
}
