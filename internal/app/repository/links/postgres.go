package links

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

func (r *PostgresRepository) ChangeUrl(ctx context.Context, url string, shortLink string) error {
	_, err := r.connPool.Exec(ctx, `
INSERT INTO links 
		(url,
		short_link)
VALUES ($1,$2);`,
		url,
		shortLink)
	if err != nil {
		return errors.Wrap(err, "failed to change url from Exec")
	}

	return nil
}

func (r *PostgresRepository) GetUrl(ctx context.Context, shortLink string) (string, error) {
	var url string

	err := r.connPool.QueryRow(ctx, `SELECT url FROM links WHERE short_link = $1`, shortLink).Scan(&url)
	if err != nil {
		return "", errors.Wrap(err, "failed to get url from QueryRow")
	}

	return url, nil
}
