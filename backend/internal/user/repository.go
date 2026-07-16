package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) Create(ctx context.Context, email, passwordHash, displayName string) (*User, error) {
	query := `
	INSERT INTO users (email, password_hash, display_name)
		VALUES ($1, $2, $3)
		RETURNING id, email, password_hash, display_name, created_at, updated_at
	`

	var u User
	err := r.pool.QueryRow(ctx, query, email, passwordHash, displayName).Scan(
		&u.ID, &u.Email, &u.PasswordHash, &u.DisplayName, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("user: failed to create: %w", err)
	}

	return &u, nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, password_hash, display_name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var u User
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&u.ID, &u.Email, &u.PasswordHash, &u.DisplayName, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("user: failed to find by email: %w", err)
	}

	return &u, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT id, email, password_hash, display_name, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var u User
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&u.ID, &u.Email, &u.PasswordHash, &u.DisplayName, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("user: failed to find by id: %w", err)
	}

	return &u, nil
}
