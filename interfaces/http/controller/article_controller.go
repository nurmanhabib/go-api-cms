package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-api-cms/app"
	"go-api-cms/infrastructure/services"
	"go-api-cms/interfaces/http/middleware"
)

type ArticleController struct {
	app *app.App
}

func NewArticleController(app *app.App) *ArticleController {
	return &ArticleController{app: app}
}

func (a *ArticleController) GetArticles(c *gin.Context) {

}

func (a *ArticleController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var newArticle services.ArticleRequest

	if err := c.ShouldBind(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user := middleware.GetUserFromContext(ctx); user != nil {
		newArticle.CreatedBy = user.ID
	}

	svc := services.NewArticleService(a.app)
	article, articleVersion, err := svc.Create(ctx, &newArticle)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"article": article,
			"version": articleVersion,
		},
	})
}

func (a *ArticleController) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var newArticle services.ArticleRequest

	slug := c.Param("slug")

	svc := services.NewArticleService(a.app)
	article, articleVersion, err := svc.GetBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Article not found",
			})
			c.Abort()
			return
		}

		c.Error(err)
		return
	}

	if err := c.ShouldBind(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user := middleware.GetUserFromContext(ctx); user != nil {
		newArticle.CreatedBy = user.ID
	}

	article, articleVersion, err = svc.Update(ctx, article.ID.String(), &newArticle)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"article": article,
			"version": articleVersion,
		},
	})
}

func (a *ArticleController) GetArticlesBySlug(c *gin.Context) {
	ctx := c.Request.Context()
	slug := c.Param("slug")

	svc := services.NewArticleService(a.app)
	article, articleVersion, err := svc.GetBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Article not found",
			})
			c.Abort()
			return
		}

		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"article": article,
			"version": articleVersion,
		},
	})
}

func (a *ArticleController) GetArticleVersionsBySlug(c *gin.Context) {
	ctx := c.Request.Context()
	slug := c.Param("slug")

	svc := services.NewArticleService(a.app)
	article, articleVersions, err := svc.GetVersionsBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Article not found",
			})
			c.Abort()
			return
		}

		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"article":  article,
			"versions": articleVersions,
		},
	})
}
