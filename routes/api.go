package routes

import (
	"github.com/gin-gonic/gin"

	"go-api-cms/app"
	"go-api-cms/interfaces/http/controller"
	"go-api-cms/interfaces/http/middleware"
)

func Api(app *app.App) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler())

	pingRoutes(router)
	userRegisterRoutes(router, app)
	articleRoutes(router, app)

	return router
}

func pingRoutes(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func userRegisterRoutes(e *gin.Engine, app *app.App) {
	c := controller.NewUserRegisterController(app)

	e.POST("/api/v1/user/register", c.Register)
}

func articleRoutes(e *gin.Engine, app *app.App) {
	c := controller.NewArticleController(app)

	e.GET("/api/v1/articles/:slug", c.GetArticlesBySlug)
}
