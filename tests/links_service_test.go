package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/BlackRRR/shortened-links/internal/app/handler"
	"github.com/BlackRRR/shortened-links/internal/app/services"
	mock_services "github.com/BlackRRR/shortened-links/internal/app/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLinksService(t *testing.T) {

	testTable := []struct {
		name                 string
		requestBody          handler.ChangeUrlRequest
		mock                 func(ctx context.Context, links *mock_services.MockLinks, url string)
		expectedResponseBody handler.ChangeUrlResponse
	}{
		{
			name: "OK",
			requestBody: handler.ChangeUrlRequest{
				Url: "https://vk.com/feed",
			},
			mock: func(ctx context.Context, links *mock_services.MockLinks, url string) {
				links.EXPECT().ChangeURL(ctx, url).Return("", nil)
			},
			expectedResponseBody: handler.ChangeUrlResponse{
				Payload: &handler.ShortenedLinkPayload{},
			},
		},
	}

	for _, testCase := range testTable {
		ctx := context.TODO()
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			links := mock_services.NewMockLinks(c)
			testCase.mock(ctx, links, testCase.requestBody.Url)

			service := &services.Services{Links: links}

			h := &handler.Handler{
				Links: service,
			}

			r := gin.New()
			r.POST("/", h.ChangeURL)

			marshal, err := json.Marshal(testCase.requestBody)
			if err != nil {
				return
			}

			body := bytes.NewReader(marshal)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", body)

			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedResponseBody, w.Body)
		})
	}
}
