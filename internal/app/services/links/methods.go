package links

import (
	"context"
	"github.com/BlackRRR/shortened-links/internal/utils"
	"github.com/pkg/errors"
)

func (s *LinksService) ChangeURL(ctx context.Context, url string) (string, error) {
	shortLink := utils.GetShortURL()

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
