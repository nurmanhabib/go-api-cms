package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-api-cms/app"
	"go-api-cms/infrastructure/services"
)

type ArticleController struct {
	app *app.App
}

func NewArticleController(app *app.App) *ArticleController {
	return &ArticleController{app: app}
}

func (a *ArticleController) GetArticles(c *gin.Context) {

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
