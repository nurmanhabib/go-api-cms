package repository

import (
	"context"

	"go-api-cms/domain/entity"
)

type ArticleRepository interface {
	FindBySlug(ctx context.Context, slug string) (*entity.Article, error)
	FindByID(ctx context.Context, articleID string) (*entity.Article, error)
	Create(ctx context.Context, article *entity.Article) error
	Update(ctx context.Context, articleID string, article *entity.Article) error
}
