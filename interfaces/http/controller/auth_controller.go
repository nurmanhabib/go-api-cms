package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api-cms/app"
	"go-api-cms/infrastructure/services"
)

type AuthController struct {
	app *app.App
}

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthController(app *app.App) *AuthController {
	return &AuthController{app: app}
}

func (ac *AuthController) Login(c *gin.Context) {
	var credentials Credentials

	err := c.ShouldBind(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	svc := services.NewUserAuthService(ac.app)
	user, errLogin := svc.Login(c, credentials.Username, credentials.Password)
	if errLogin != nil {
		_ = c.Error(errLogin)
		return
	}

	tokenSvc := services.NewAuthService(ac.app)
	token, errToken := tokenSvc.GenerateToken(user)
	if errToken != nil {
		_ = c.Error(errToken)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": token}})
}
