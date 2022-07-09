package repository

import (
	"context"
	"github.com/BlackRRR/shortened-links/internal/app/repository/links"
	"github.com/BlackRRR/shortened-links/internal/app/repository/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

//Linker any database that satisfies the methods can be used
type Linker interface {
	ChangeUrl(ctx context.Context, url string, shortLink string) error
	GetUrl(ctx context.Context, shortLink string) (string, error)
}

type Repositories struct {
	Postgres *postgres.PostgresRepository
	Links    *links.Links
}

func InitRepositories(ctx context.Context, connPool *pgxpool.Pool) (*Repositories, error) {
	//Init postgres repository
	postgresRepository, err := postgres.InitPostgresRepository(ctx, connPool)
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
