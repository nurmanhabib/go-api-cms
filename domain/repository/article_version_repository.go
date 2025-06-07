package repository

import (
	"context"

	"go-api-cms/domain/entity"
)

type ArticleVersionRepository interface {
	FindByArticleID(ctx context.Context, articleID string) ([]*entity.ArticleVersion, error)
	FindByArticleIDAndVersion(ctx context.Context, articleID string, version int64) (*entity.ArticleVersion, error)
	FindByID(ctx context.Context, articleVersionID string) (*entity.ArticleVersion, error)
	Create(ctx context.Context, articleVersion *entity.ArticleVersion) error
	Update(ctx context.Context, articleVersionID string, articleVersion *entity.ArticleVersion) error
	NextVersion(ctx context.Context, articleID string) (int64, error)
}
