package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

func InitDataBase(ctx context.Context, cfg *pgxpool.Config) (*pgxpool.Pool, error) {
	//ConnectConfig creates a new Pool and immediately establishes one connection.
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create conn pool")
	}

	return pool, nil
}
