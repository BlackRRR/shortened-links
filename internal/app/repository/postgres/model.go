package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type PostgresRepository struct {
	connPool *pgxpool.Pool
}

func InitPostgresRepository(ctx context.Context, db *pgxpool.Pool) (*PostgresRepository, error) {
	PostgresRep := &PostgresRepository{connPool: db}

	rows, err := PostgresRep.connPool.Query(ctx, `
CREATE TABLE IF NOT EXISTS links(
	url text, 
	short_link text UNIQUE
);`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create links table")
	}
	rows.Close()

	return PostgresRep, nil
}
