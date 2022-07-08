package server

import (
	"context"
	"github.com/BlackRRR/shortened-Links/internal/models"
)

func (s *Server) ChangeURL(ctx context.Context, request *ChangeUrlRequest) (*ChangeUrlResponse, error) {
	shortLink, err := s.Links.ChangeURL(ctx, request.Url)
	if err != nil {
		return &ChangeUrlResponse{
			Result:  models.ResultERR,
			Payload: nil,
			Error:   models.NewLinksError(err),
		}, err
	}

	return &ChangeUrlResponse{
		Result:  models.ResultOK,
		Payload: &ShortenedLinkPayload{shortLink},
		Error:   nil,
	}, nil
}

func (s *Server) GetURL(ctx context.Context, request *GetUrlRequest) (*GetUrlResponse, error) {
	url, err := s.Links.GetURL(ctx, request.Link)
	if err != nil {
		return &GetUrlResponse{
			Result:  models.ResultERR,
			Payload: nil,
			Error:   models.NewLinksError(err),
		}, err
	}

	return &GetUrlResponse{
		Result:  models.ResultOK,
		Payload: &UrlPayload{Url: url},
		Error:   nil,
	}, nil
}
