package links

import (
	"context"
)

type MapRepository struct {
	link map[string]string
}

func NewLinksRepository() *MapRepository {
	linksRep := &MapRepository{link: make(map[string]string)}

	return linksRep
}

func (r *MapRepository) ChangeUrl(_ context.Context, url string, shortLink string) (string, error) {
	r.link[shortLink] = url

	return shortLink, nil
}

func (r *MapRepository) GetUrl(_ context.Context, shortLink string) (string, error) {
	return r.link[shortLink], nil
}
