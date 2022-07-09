package links

import "github.com/BlackRRR/shortened-links/internal/app/repository"

type LinksService struct {
	repository.Linker
}

func InitLinksService(links repository.Linker) *LinksService {
	return &LinksService{links}
}
