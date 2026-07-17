package refreshtoken

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) Store(ctx context.Context, userID, tokenHash string, ttl time.Duration) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token_hash, expires_at)
		VALUES ($1, $2, $3)
	`

	expiresAt := time.Now().Add(ttl)

	_, err := r.pool.Exec(ctx, query, userID, tokenHash, expiresAt)
	if err != nil {
		return fmt.Errorf("refreshtoken: failed to store: %w", err)
	}

	return nil
}

func (r *Repository) FindUserIDByHash(ctx context.Context, tokenHash string) (string, error) {
	query := `
		SELECT user_id
		FROM refresh_tokens
		WHERE token_hash = $1 AND expires_at > now()
	`

	var userID string
	err := r.pool.QueryRow(ctx, query, tokenHash).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("refreshtoken: failed to find: %w", err)
	}

	return userID, nil
}

func (r *Repository) DeleteByHash(ctx context.Context, tokenHash string) error {
	query := `DELETE FROM refresh_tokens WHERE token_hash = $1`

	_, err := r.pool.Exec(ctx, query, tokenHash)
	if err != nil {
		return fmt.Errorf("refreshtoken: failed to delete: %w", err)
	}

	return nil
}
