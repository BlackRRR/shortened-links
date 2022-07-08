package links

import "context"

func (l *Links) ChangeUrl(_ context.Context, url string, shortLink string) error {
	l.link[shortLink] = url
	return nil
}

func (l *Links) GetUrl(_ context.Context, shortLink string) (string, error) {
	return l.link[shortLink], nil
}
