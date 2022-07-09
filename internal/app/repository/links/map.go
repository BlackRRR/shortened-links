package links

import "context"

type Links struct {
	link map[string]string
}

func InitLinksRepository() *Links {
	linksRep := &Links{link: make(map[string]string)}

	return linksRep
}

func (l *Links) ChangeUrl(_ context.Context, url string, shortLink string) error {
	l.link[shortLink] = url
	return nil
}

func (l *Links) GetUrl(_ context.Context, shortLink string) (string, error) {
	return l.link[shortLink], nil
}
