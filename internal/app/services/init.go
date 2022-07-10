package services

import (
	"github.com/BlackRRR/shortened-links/internal/app/repository"
	"github.com/BlackRRR/shortened-links/internal/app/services/links"
)

type Services struct {
	*links.LinksService
}

func InitServices(repository *repository.Repositories) *Services {
	//Init Map service
	//We have Map and Postgres repositories, and we change them simply by passing different parameters to the function InitLinksService
	return &Services{links.InitLinksService(repository.Map)}
}
