package services

import (
	"context"
	"github.com/BlackRRR/shortened-links/internal/app/repository"
	"github.com/BlackRRR/shortened-links/internal/app/services/links"
)

//go:generate mockgen -source=init.go -destination=mocks/mock.go

//Links any service that satisfies the methods can be used

type Links interface {
	ChangeURL(_ context.Context, url string) (string, error)
	GetURL(_ context.Context, shortLink string) (string, error)
}

type Services struct {
	Links
}

func InitServices(repository *repository.Repositories) *Services {
	//Init Links service
	return &Services{Links: links.InitLinksService(repository.Links)}
}
