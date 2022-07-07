package links

import "context"

func (l *Links) ChangeUrl(_ context.Context, url string, link string) error {
	l.link[link] = url
	return nil
}

func (l *Links) GetUrl(_ context.Context, link string) (string, error) {
	return l.link[link], nil
}
