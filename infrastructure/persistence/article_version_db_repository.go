package persistence

import (
	"context"

	"gorm.io/gorm"

	"go-api-cms/domain/entity"
)

type ArticleVersionDBRepository struct {
	db *gorm.DB
}

func NewArticleVersionDBRepository(db *gorm.DB) *ArticleVersionDBRepository {
	return &ArticleVersionDBRepository{db: db}
}

func (a *ArticleVersionDBRepository) FindByArticleIDAndVersion(ctx context.Context, articleID string, version int64) (*entity.ArticleVersion, error) {
	var articleVersion entity.ArticleVersion

	err := a.db.WithContext(ctx).First(&articleVersion, "article_id = ? AND version = ?", articleID, version).Error
	if err != nil {
		return nil, err
	}

	return &articleVersion, nil
}

func (a *ArticleVersionDBRepository) FindByID(ctx context.Context, articleVersionID string) (*entity.ArticleVersion, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleVersionDBRepository) Create(ctx context.Context, articleVersion *entity.ArticleVersion) error {
	return a.db.Create(articleVersion).Error
}

func (a *ArticleVersionDBRepository) Update(ctx context.Context, articleVersionID string, articleVersion *entity.ArticleVersion) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleVersionDBRepository) NextVersion(ctx context.Context, articleID string) (int64, error) {
	var nextVersion int64

	err := a.db.WithContext(ctx).
		Raw(`
			SELECT COALESCE(MAX(version), 0) + 1
			FROM article_versions
			WHERE article_id = ?
		`, articleID).
		Scan(&nextVersion).Error
	if err != nil {
		return 0, err
	}

	return nextVersion, nil
}
