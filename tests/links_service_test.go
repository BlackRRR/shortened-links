package tests

import (
	"github.com/BlackRRR/shortened-Links/internal/app/server"
	"github.com/BlackRRR/shortened-Links/internal/models"
	"testing"
)

func TestLinksService(t *testing.T) {
	testTable := []struct {
		name                 string
		requestBody          server.ChangeUrlRequest
		expectedResponseBody server.ChangeUrlResponse
	}{
		{
			name: "OK",
			requestBody: server.ChangeUrlRequest{
				Url: "",
			},
			expectedResponseBody: server.ChangeUrlResponse{
				Result:  models.ResultOK,
				Payload: &server.ShortenedLinkPayload{},
				Error:   nil,
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

		})
	}
}
