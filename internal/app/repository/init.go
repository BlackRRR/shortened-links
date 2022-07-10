package repository

import (
	"context"
	"github.com/BlackRRR/shortened-links/internal/app/repository/links"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Repositories struct {
	Postgres *links.PostgresRepository
	Map      *links.MapRepository
}

func InitRepositories(ctx context.Context, connPool *pgxpool.Pool) (*Repositories, error) {
	//Init postgres repository
	postgresRepository, err := links.NewPostgresRepository(ctx, connPool)
	if err != nil {
		return nil, errors.Wrap(err, "failed init postgres repository")
	}

	//Init links repository
	linksRepository := links.NewLinksRepository()

	init := &Repositories{
		Postgres: postgresRepository,
		Map:      linksRepository,
	}

	return init, nil
}
