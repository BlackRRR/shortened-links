package links

import (
	"context"
	"github.com/BlackRRR/shortened-Links/internal/app/repository"
)

type Ser interface {
	ChangeURL(ctx context.Context, url string) (string, error)
	GetURL(ctx context.Context, link string) (string, error)
}

type LinksService struct {
	links repository.Linker
}

func InitLinksService(links repository.Linker) Ser {
	return &LinksService{links: links}
}
