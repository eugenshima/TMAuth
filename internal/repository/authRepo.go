// Package repository contains methods for interacting with database
package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/TMAuth/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

// ProfileRepository struct
type ProfileRepository struct {
	pool *pgxpool.Pool
}

// NewProfileRepository creates a new ProfileRepository
func NewProfileRepository(pool *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{pool: pool}
}

func (db *ProfileRepository) GetProfileByLogin(ctx context.Context, login string) (*model.FullAuthModel, error) {
	profile := &model.FullAuthModel{}
	err := db.pool.QueryRow(context.Background(), "select id, login, password from profile.profile WHERE login = $1", login).Scan(&profile.ID, &profile.Login, &profile.Password)
	if err != nil {
		return nil, fmt.Errorf("GetProfileByLogin: %w", err)
	}
	return profile, nil
}
