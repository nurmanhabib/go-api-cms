package entity

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID                uuid.UUID `json:"id"`
	PermissionGroupID uuid.UUID `json:"permission_group_id"`
	Name              string    `json:"name"`
	Label             string    `json:"label"`
	Description       string    `json:"description"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
