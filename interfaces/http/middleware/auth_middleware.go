package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-api-cms/app"
	"go-api-cms/domain/entity"
	"go-api-cms/infrastructure/services"
)

type ctxKey string

const ctxKeyUser ctxKey = "user"

func WithUser(ctx context.Context, user *entity.User) context.Context {
	return context.WithValue(ctx, ctxKeyUser, user)
}

func GetUserFromContext(ctx context.Context) *entity.User {
	if user, ok := ctx.Value(ctxKeyUser).(*entity.User); ok {
		return user
	}
	return nil
}

func AuthMiddleware(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		tokenStr := parts[1]
		ctx := c.Request.Context()

		svc := services.NewAuthService(app)
		loggedInUser, err := svc.ParseToken(ctx, tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(ctx, ctxKeyUser, loggedInUser))

		c.Next()
	}
}
