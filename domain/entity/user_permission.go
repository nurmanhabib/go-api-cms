package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserPermission struct {
	UserID       uuid.UUID
	PermissionID uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
