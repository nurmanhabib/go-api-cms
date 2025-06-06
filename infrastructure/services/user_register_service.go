package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"go-api-cms/app"
	"go-api-cms/domain/entity"
	"go-api-cms/infrastructure/exception"
)

type UserRegisterService struct {
	app *app.App
}

func NewUserRegisterService(app *app.App) *UserRegisterService {
	return &UserRegisterService{app: app}
}

func (ur *UserRegisterService) Registration(ctx context.Context, user *entity.User, assignedRoles []string) (err error) {
	txRepo, tx := ur.app.Repo.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err = txRepo.UserRepo.Create(ctx, user); err != nil {
		if strings.Contains(err.Error(), "users_username_key") {
			return exception.NewUnprocessableEntity([]exception.FieldViolation{
				{
					Field:       "username",
					Description: "Already exists",
				},
			})
		}

		if strings.Contains(err.Error(), "users_email_key") {
			return exception.NewUnprocessableEntity([]exception.FieldViolation{
				{
					Field:       "email",
					Description: "Already exists",
				},
			})
		}

		return fmt.Errorf("create user error: %w", err)
	}

	var roleIDs []uuid.UUID

	roles, errRoles := txRepo.RoleRepo.FindByNames(ctx, assignedRoles)
	if errRoles != nil {
		return errRoles
	}

	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID)
	}

	if err = txRepo.UserRepo.AssignRoles(ctx, user.ID, roleIDs); err != nil {
		return fmt.Errorf("assign user error: %w", err)
	}

	return nil
}
