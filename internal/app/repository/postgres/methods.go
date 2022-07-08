package postgres

import (
	"context"
	"github.com/pkg/errors"
)

func (r *PostgresRepository) ChangeUrl(ctx context.Context, url string, link string) error {
	_, err := r.connPool.Exec(ctx, `
INSERT INTO links 
		(url,
		link)
VALUES ($1,$2);`,
		url,
		link)
	if err != nil {
		return errors.Wrap(err, "failed to change url from Exec")
	}

	return nil
}

func (r *PostgresRepository) GetUrl(ctx context.Context, link string) (string, error) {
	var URL string

	err := r.connPool.QueryRow(ctx, `SELECT url FROM links WHERE link = $1`, link).Scan(&URL)
	if err != nil {
		return "", errors.Wrap(err, "failed to get url from QueryRow")
	}

	return URL, nil
}
