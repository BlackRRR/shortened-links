package links

import (
	"context"
	"github.com/pkg/errors"
	"math/rand"
)

var (
	symbols       = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789_")
	symbolsLength = 10
)

type LinksService struct {
	Repository
}

func InitLinksService(links Repository) *LinksService {
	return &LinksService{links}
}

func (s *LinksService) ChangeURL(ctx context.Context, url string) (string, error) {
	shortLink := GetShortURL()

	shortLink, err := s.Repository.ChangeUrl(ctx, url, shortLink)
	if err != nil {
		return "", errors.Wrap(err, "service: failed to change url")
	}

	return shortLink, nil
}

func (s *LinksService) GetURL(ctx context.Context, shortLink string) (string, error) {
	url, err := s.Repository.GetUrl(ctx, shortLink)
	if err != nil {
		return "", errors.Wrap(err, "service: failed to get url")
	}

	if url == "" {
		return "", nil
	}

	return url, nil
}

func GetShortURL() string {
	key := make([]rune, symbolsLength)

	for i := range key {
		key[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(key)
}
