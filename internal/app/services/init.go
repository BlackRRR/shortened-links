package services

import (
	"github.com/BlackRRR/shortened-Links/internal/app/repository"
)

type Services struct {
	Links *LinksService
}

func InitServices(repository *repository.Repositories) *Services {
	//Init Links service
	return &Services{Links: InitLinksService(repository.Links)}
}
