package services

import (
	"github.com/BlackRRR/shortened-Links/internal/app/repository"
	"github.com/BlackRRR/shortened-Links/internal/app/services/links"
)

type Services struct {
	Links links.Ser
}

func InitServices(repository *repository.Repositories) *Services {
	//Init Links service
	linksSrv := links.InitLinksService(repository.Links)

	return &Services{Links: linksSrv}
}
