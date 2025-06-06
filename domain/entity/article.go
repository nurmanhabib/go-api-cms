package entity

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID               uuid.UUID `json:"id"`
	Slug             string    `json:"slug"`
	CurrentVersionID uuid.UUID `json:"current_version_id"`
	IsDeleted        bool      `json:"-"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
