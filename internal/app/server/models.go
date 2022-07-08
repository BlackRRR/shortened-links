package server

import "github.com/BlackRRR/shortened-Links/internal/models"

//////////////////////////
//Change Url Request
//////////////////////////

type ChangeUrlRequest struct {
	Url string `json:"url"`
}

type ChangeUrlResponse struct {
	Result  models.Result         `json:"result"`
	Payload *ShortenedLinkPayload `json:"payload"`
	Error   *models.ServerError   `json:"error"`
}

type ShortenedLinkPayload struct {
	Link string `json:"link"`
}

//////////////////////////
//Get Url Request
//////////////////////////

type GetUrlRequest struct {
	Link string `json:"link"`
}

type GetUrlResponse struct {
	Result  models.Result       `json:"result"`
	Payload *UrlPayload         `json:"payload"`
	Error   *models.ServerError `json:"error"`
}

type UrlPayload struct {
	Url string `json:"url"`
}
