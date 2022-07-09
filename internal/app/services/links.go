package services

import (
	"context"
	"github.com/pkg/errors"
	"math/rand"
)

var (
	symbols       = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789_")
	symbolsLength = 10
)

//go:generate mockgen -source=links.go -destination=mocks/mock.go

type Linker interface {
	ChangeUrl(ctx context.Context, url string, shortLink string) error
	GetUrl(ctx context.Context, shortLink string) (string, error)
}

type LinksService struct {
	Linker
}

func InitLinksService(links Linker) *LinksService {
	return &LinksService{links}
}

func (s *LinksService) ChangeURL(ctx context.Context, url string) (string, error) {
	shortLink := GetShortURL()

	err := s.Linker.ChangeUrl(ctx, url, shortLink)
	if err != nil {
		return "", errors.Wrap(err, "service: failed to change url")
	}

	return shortLink, nil
}

func (s *LinksService) GetURL(ctx context.Context, shortLink string) (string, error) {
	url, err := s.Linker.GetUrl(ctx, shortLink)
	if err != nil {
		return "", errors.Wrap(err, "service: failed to get url")
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
