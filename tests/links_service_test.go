package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/BlackRRR/shortened-links/internal/app/handler"
	"github.com/BlackRRR/shortened-links/internal/app/repository"
	"github.com/BlackRRR/shortened-links/internal/app/repository/links"
	"github.com/BlackRRR/shortened-links/internal/app/services"
	links2 "github.com/BlackRRR/shortened-links/internal/app/services/links"
	"github.com/magiconair/properties/assert"
	"io"
	"net/http"
	"testing"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func TestGetShortLink(t *testing.T) {
	shortLink := links2.GetShortURL()

	if len(shortLink) != 10 {
		t.Error("failed to get short url")
	}

	t.Log(shortLink)
}

func TestChangeUrlService(t *testing.T) {
	ctx := context.TODO()

	answer, err := services.InitServices(&repository.Repositories{
		Map: links.NewLinksRepository(),
	}).LinksService.ChangeURL(ctx, "")
	if err != nil {
		t.Errorf("failed test server: %v", err)
	}

	if answer == "" {
		t.Error("no short links in result set")
	}

	t.Log(answer)
}

func TestGetUrlService(t *testing.T) {
	ctx := context.TODO()

	answer, err := services.InitServices(&repository.Repositories{
		Map: links.NewLinksRepository(),
	}).LinksService.GetUrl(ctx, "Tjqelfsb8o")
	if err != nil {
		t.Errorf("failed test server: %v", err)
	}

	if answer == "" {
		t.Error("no urls in result set")
	}

	t.Log(answer)
}

func TestChangeUrlLinksHttp(t *testing.T) {
	linkResponse := handler.ChangeUrlResponse{
		Payload: &handler.ShortenedLinkPayload{},
	}

	marshal, err := json.Marshal(&handler.ChangeUrlRequest{
		Url: "https://mail.google.com/mail/u/0/#inbox",
	})
	if err != nil {
		t.Errorf("failed marshal payment request %v", err)
	}

	r := bytes.NewReader(marshal)

	c := NewClient("http://localhost:8011/")
	res, err := c.NewRequest(r, http.MethodPost)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = res.Decode(&linkResponse)
	if err != nil {
		t.Errorf("failed to decode response %v", err)
	}

	if linkResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	t.Log(linkResponse.Payload.Link)
}

func TestGetLinkLinksHttp(t *testing.T) {
	linkResponse := handler.GetUrlResponse{
		Payload: &handler.UrlPayload{Url: ""},
	}

	expected := handler.GetUrlResponse{
		Payload: &handler.UrlPayload{Url: "https://mail.google.com/mail/u/0/#inbox"},
	}

	c := NewClient("http://localhost:8011/Tjqelfsb8o")
	res, err := c.NewRequest(nil, http.MethodGet)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = res.Decode(&linkResponse)
	if err != nil {
		t.Errorf("failed to decode response %v", err)
	}

	if linkResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	assert.Equal(t, expected, linkResponse)

	t.Log(linkResponse.Payload.Url)
}

func (c Client) NewRequest(r io.Reader, method string) (*json.Decoder, error) {
	client := http.Client{}
	request, err := http.NewRequest(method, c.url, r)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	decode := json.NewDecoder(resp.Body)

	return decode, nil
}
