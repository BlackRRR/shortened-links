package links

import (
	"context"
	"github.com/pkg/errors"
)

func (s *LinksService) ChangeURL(ctx context.Context, url string) (string, error) {
	err := s.links.ChangeUrl(ctx, url, "") //TODO: CreateLink
	if err != nil {
		return "", errors.Wrap(err, "service: failed to change url")
	}

	return "", nil

}

func (s *LinksService) GetURL(ctx context.Context, link string) (string, error) {
	url, err := s.links.GetUrl(ctx, link)
	if err != nil {
		return "", errors.Wrap(err, "service: failed to get url")
	}

	return url, nil
}
