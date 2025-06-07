package services

import (
	"context"

	"github.com/google/uuid"

	"go-api-cms/app"
	"go-api-cms/domain/entity"
)

type ArticleService struct {
	app *app.App
}

type ArticleRequest struct {
	Slug      string    `json:"slug,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Status    string    `json:"status,omitempty"`
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
}

func NewArticleService(app *app.App) *ArticleService {
	return &ArticleService{app: app}
}

func (a *ArticleService) GetBySlug(ctx context.Context, slug string) (*entity.Article, *entity.ArticleVersion, error) {
	article, err := a.app.Repo.ArticleRepo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, nil, err
	}

	articleVersion, err := a.app.Repo.ArticleVersionRepo.FindByID(ctx, article.CurrentVersionID.String())
	if err != nil {
		return article, nil, err
	}

	return article, articleVersion, nil
}

func (a *ArticleService) GetVersionsBySlug(ctx context.Context, slug string) (*entity.Article, []*entity.ArticleVersion, error) {
	article, err := a.app.Repo.ArticleRepo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, nil, err
	}

	articleVersions, err := a.app.Repo.ArticleVersionRepo.FindByArticleID(ctx, article.ID.String())
	if err != nil {
		return article, nil, err
	}

	return article, articleVersions, nil
}

func (a *ArticleService) Create(ctx context.Context, newArticle *ArticleRequest) (article *entity.Article, articleVersion *entity.ArticleVersion, err error) {
	txRepo, tx := a.app.Repo.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	newEntityArticle := &entity.Article{
		ID:               uuid.New(),
		Slug:             newArticle.Slug,
		CurrentVersionID: uuid.New(),
	}

	newEntityArticleVersion := &entity.ArticleVersion{
		ID:        newEntityArticle.CurrentVersionID,
		ArticleID: newEntityArticle.ID,
		Version:   1,
		Title:     newArticle.Title,
		Content:   newArticle.Content,
		Status:    newArticle.Status,
		CreatedBy: newArticle.CreatedBy,
	}

	errArticle := txRepo.ArticleRepo.Create(ctx, newEntityArticle)
	if errArticle != nil {
		return nil, nil, errArticle
	}

	errVersion := txRepo.ArticleVersionRepo.Create(ctx, newEntityArticleVersion)
	if errVersion != nil {
		return newEntityArticle, nil, errVersion
	}

	return newEntityArticle, newEntityArticleVersion, nil
}

func (a *ArticleService) Update(ctx context.Context, articleID string, editArticle *ArticleRequest) (article *entity.Article, articleVersion *entity.ArticleVersion, err error) {
	txRepo, tx := a.app.Repo.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	existsArticle, err := a.app.Repo.ArticleRepo.FindByID(ctx, articleID)
	if err != nil {
		return nil, nil, err
	}

	nextVersion, errNextVersion := a.app.Repo.ArticleVersionRepo.NextVersion(ctx, articleID)
	if errNextVersion != nil {
		return existsArticle, nil, errNextVersion
	}

	newEntityArticleVersion := &entity.ArticleVersion{
		ID:        uuid.New(),
		ArticleID: uuid.MustParse(articleID),
		Version:   nextVersion,
		Title:     editArticle.Title,
		Content:   editArticle.Content,
		Status:    editArticle.Status,
		CreatedBy: editArticle.CreatedBy,
	}

	errVersion := txRepo.ArticleVersionRepo.Create(ctx, newEntityArticleVersion)
	if errVersion != nil {
		return existsArticle, nil, errVersion
	}

	if editArticle.Status == "published" {
		existsArticle.CurrentVersionID = newEntityArticleVersion.ID
		errUpdate := txRepo.ArticleRepo.Update(ctx, existsArticle.ID.String(), existsArticle)
		if errUpdate != nil {
			return existsArticle, newEntityArticleVersion, errUpdate
		}
	}

	return existsArticle, newEntityArticleVersion, nil
}

func (a *ArticleService) RollbackVersion(ctx context.Context, articleID string, version int64) (article *entity.Article, articleVersion *entity.ArticleVersion, err error) {
	txRepo, tx := a.app.Repo.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	existsArticle, err := txRepo.ArticleRepo.FindByID(ctx, articleID)
	if err != nil {
		return nil, nil, err
	}

	existsVersion, err := txRepo.ArticleVersionRepo.FindByArticleIDAndVersion(ctx, articleID, version)
	if err != nil {
		return existsArticle, nil, err
	}

	existsArticle.CurrentVersionID = existsVersion.ID

	errUpdate := txRepo.ArticleRepo.Update(ctx, articleID, existsArticle)
	if errUpdate != nil {
		return existsArticle, nil, errUpdate
	}

	return existsArticle, existsVersion, nil
}
