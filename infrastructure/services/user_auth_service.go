package services

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"go-api-cms/app"
	"go-api-cms/domain/entity"
	"go-api-cms/infrastructure/exception"
)

type UserAuthService struct {
	app *app.App
}

func NewUserAuthService(app *app.App) *UserAuthService {
	return &UserAuthService{app: app}
}

func (svc *UserAuthService) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	user, err := svc.app.Repo.UserRepo.FindByUsername(ctx, username)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, exception.NewUnprocessableEntity([]exception.FieldViolation{
				{
					Field:       "username",
					Description: "User not found",
				},
			})

		default:
			return nil, err
		}
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if errCompare != nil {
		return nil, exception.NewUnprocessableEntity([]exception.FieldViolation{
			{
				Field:       "username",
				Description: "Credentials are not correct",
			},
		})
	}

	return user, nil
}
