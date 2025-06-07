package persistence

import (
	"context"

	"gorm.io/gorm"

	"go-api-cms/domain/entity"
)

type ArticleDBRepository struct {
	db *gorm.DB
}

func NewArticleDBRepository(db *gorm.DB) *ArticleDBRepository {
	return &ArticleDBRepository{db: db}
}

func (a *ArticleDBRepository) FindBySlug(ctx context.Context, slug string) (*entity.Article, error) {
	var article entity.Article
	err := a.db.WithContext(ctx).First(&article, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *ArticleDBRepository) FindByID(ctx context.Context, articleID string) (*entity.Article, error) {
	var article entity.Article
	err := a.db.WithContext(ctx).First(&article, "id = ?", articleID).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *ArticleDBRepository) Create(ctx context.Context, article *entity.Article) error {
	return a.db.Create(article).Error
}

func (a *ArticleDBRepository) Update(ctx context.Context, articleID string, article *entity.Article) error {
	return a.db.WithContext(ctx).Where("id = ?", articleID).Updates(&article).Error
}
