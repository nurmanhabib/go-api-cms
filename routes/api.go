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

	authRoutes(router, app)
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

func authRoutes(e *gin.Engine, app *app.App) {
	c := controller.NewAuthController(app)

	e.POST("/api/v1/auth/token", c.Login)
}

func userRegisterRoutes(e *gin.Engine, app *app.App) {
	c := controller.NewUserRegisterController(app)

	e.POST("/api/v1/user/register", c.Register)
}

func articleRoutes(e *gin.Engine, app *app.App) {
	c := controller.NewArticleController(app)

	g := e.Group("/", middleware.AuthMiddleware(app))
	g.GET("/api/v1/articles/:slug", c.GetArticlesBySlug)
}
