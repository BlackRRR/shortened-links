package server

import (
	"context"
	"encoding/json"
	"github.com/BlackRRR/shortened-Links/internal/app/services"
	"github.com/BlackRRR/shortened-Links/internal/app/services/links"
	"github.com/gorilla/mux"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

type Server struct {
	Links *links.LinksService
}

// Server Init

func NewServer(service *services.Services) *Server {
	return &Server{Links: service.Links}

}

func MakeHTTPHandler(s *Server) http.Handler {
	h := mux.NewRouter()

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(NewEncodeError()),
	}

	initLinksHandlers(h, s, opts)

	return h
}

func newEncodeResponse() httptransport.EncodeResponseFunc {
	return func(_ context.Context, w http.ResponseWriter, response interface{}) error {
		w.WriteHeader(http.StatusOK)

		return json.NewEncoder(w).Encode(response)
	}
}

func NewEncodeError() httptransport.ErrorEncoder {
	return func(_ context.Context, err error, w http.ResponseWriter) {

		w.WriteHeader(http.StatusInternalServerError)

		_ = json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}
}
