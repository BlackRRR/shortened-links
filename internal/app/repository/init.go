package repository

import (
	"context"
	"github.com/BlackRRR/shortened-Links/internal/app/repository/links"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Repositories struct {
	Postgres *links.PostgresRepository
	Links    *links.Links
}

func InitRepositories(ctx context.Context, connPool *pgxpool.Pool) (*Repositories, error) {
	//Init postgres repository
	postgresRepository, err := links.InitPostgresRepository(ctx, connPool)
	if err != nil {
		return nil, errors.Wrap(err, "failed init Links repository")
	}

	linksRepository := links.InitLinksRepository()

	init := &Repositories{
		Postgres: postgresRepository,
		Links:    linksRepository,
	}

	return init, nil
}
