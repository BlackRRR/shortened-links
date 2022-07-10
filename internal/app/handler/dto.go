package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewLinksError(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

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
