package server

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
)

type request struct {
	weight  int
	payload []byte
}

func initLinksHandlers(h *mux.Router, s *Server, opts []httptransport.ServerOption) {
	ChangeUrlEndpoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(*ChangeUrlRequest)
		return s.ChangeURL(ctx, r)
	}

	GetUrlEndpoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(*GetUrlRequest)
		return s.GetURL(ctx, r)
	}

	h.Methods("POST").Path("/change-url").Handler(httptransport.NewServer(
		ChangeUrlEndpoint,
		decodeChangeUrlRequest,
		newEncodeResponse(),
		opts...))

	h.Methods("GET").Path("/{link}").Handler(httptransport.NewServer(
		GetUrlEndpoint,
		decodeGetUrlRequest,
		newEncodeResponse(),
		opts...))

}

func decodeChangeUrlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	incomeReq, err := readRequestBody(r)
	if err != nil {
		log.Printf("failed read http request body: %s", err.Error())
		return &ChangeUrlRequest{}, nil
	}

	req := &ChangeUrlRequest{}
	if err := json.Unmarshal(incomeReq.payload, &req); err != nil {
		return nil, errors.Wrap(err, "unmarshal request")
	}

	return req, nil
}

func decodeGetUrlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	link := mux.Vars(r)["link"]

	req := &GetUrlRequest{
		Link: link,
	}

	return req, nil
}

func readRequestBody(r *http.Request) (*request, error) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read from body")
	}

	return &request{
		weight:  len(req),
		payload: req,
	}, nil
}
