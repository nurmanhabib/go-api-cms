package entity

import (
	"time"

	"github.com/google/uuid"
)

type ArticleVersion struct {
	ID        uuid.UUID `json:"id"`
	ArticleID uuid.UUID `json:"article_id"`
	Version   int64     `json:"version"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedBy uuid.UUID `json:"created_by"`
	IsDeleted bool      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
