package handler

//////////////////////////
//Change Url Request
//////////////////////////

type ChangeUrlRequest struct {
	Url string `json:"url"`
}

type ChangeUrlResponse struct {
	Payload *ShortenedLinkPayload `json:"payload"`
}

type ShortenedLinkPayload struct {
	Link string `json:"link"`
}

//////////////////////////
//Get Url Response
//////////////////////////

type GetUrlResponse struct {
	Payload *UrlPayload `json:"payload"`
}

type UrlPayload struct {
	Url string `json:"url"`
}
