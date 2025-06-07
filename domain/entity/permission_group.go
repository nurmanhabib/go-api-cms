package entity

import (
	"time"

	"github.com/google/uuid"
)

type PermissionGroup struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
