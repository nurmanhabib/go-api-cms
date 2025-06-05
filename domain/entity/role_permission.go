package entity

import (
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	RoleID       uuid.UUID
	PermissionID uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
