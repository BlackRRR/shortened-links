package links

import (
	"github.com/BlackRRR/shortened-Links/internal/app/repository"
)

type LinksService struct {
	linker repository.Linker
}

func InitLinksService(links repository.Linker) *LinksService {
	return &LinksService{linker: links}
}
