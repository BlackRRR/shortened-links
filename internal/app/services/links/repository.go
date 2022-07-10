package links

import "context"

//all databases that implement this interface can be used

type Repository interface {
	ChangeUrl(ctx context.Context, url string, shortLink string) (string, error)
	GetUrl(ctx context.Context, shortLink string) (string, error)
}
