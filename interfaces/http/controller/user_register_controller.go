package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"go-api-cms/app"
	"go-api-cms/domain/entity"
	"go-api-cms/infrastructure/exception"
	"go-api-cms/infrastructure/services"
)

type UserRegisterController struct {
	app *app.App
}

type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required,max=10,min=1"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,max=20,min=1,min=5"`
	Password string `json:"password" validate:"required,max=20,min=6"`
	Role     string `json:"role" validate:"required,max=20,min=1,max=5"`
}

func NewUserRegisterController(app *app.App) *UserRegisterController {
	return &UserRegisterController{app: app}
}

func (u *UserRegisterController) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var r UserRegisterRequest

	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := validator.New()
	if err := v.Struct(&r); err != nil {
		var validationErrors validator.ValidationErrors
		var violations []exception.FieldViolation

		errors.As(err, &validationErrors)

		for _, validationError := range validationErrors {
			violations = append(violations, exception.FieldViolation{
				Field:       validationError.Field(),
				Description: fmt.Sprintf("%s", validationError.Tag()),
			})
		}

		_ = c.Error(exception.NewUnprocessableEntity(violations))
		return
	}

	user := &entity.User{
		ID:       uuid.New(),
		Name:     r.Name,
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		IsActive: true,
	}

	svc := services.NewUserRegisterService(u.app)
	err := svc.Registration(ctx, user, []string{r.Role})
	if err != nil {
		var svcException *exception.ServiceError
		if errors.As(err, &svcException) {
			svcException.Gin(c)
			return
		}

		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
