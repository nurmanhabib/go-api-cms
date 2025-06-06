package seeder

import (
	"context"

	"gorm.io/gorm"
)

type Seeder interface {
	Seed(ctx context.Context, db *gorm.DB) error
}

type Registry struct {
	seeders []Seeder
}

func NewRegistry(seeders ...Seeder) *Registry {
	return &Registry{
		seeders: seeders,
	}
}

func (r *Registry) Run(ctx context.Context, db *gorm.DB) error {
	for _, seeder := range r.seeders {
		if err := seeder.Seed(ctx, db); err != nil {
			return err
		}
	}

	return nil
}
