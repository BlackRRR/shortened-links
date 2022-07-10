package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) ChangeURL(c *gin.Context) {
	var req ChangeUrlRequest

	if err := c.BindJSON(&req); err != nil {
		NewLinksError(c, http.StatusBadRequest, "invalid req body")
		return
	}

	if !strings.Contains(req.Url, "https://") && !strings.Contains(req.Url, "http://") {
		NewLinksError(c, http.StatusBadRequest, "invalid req body")
		return
	}

	if req.Url == "" {
		NewLinksError(c, http.StatusBadRequest, "invalid req body")
		return
	}

	link, err := h.Services.ChangeURL(c, req.Url)
	if err != nil {
		NewLinksError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ChangeUrlResponse{
		Payload: &ShortenedLinkPayload{
			Link: link,
		},
	})
}

func (h *Handler) GetURL(c *gin.Context) {
	link := c.Param("link")

	if link == "" {
		NewLinksError(c, http.StatusInternalServerError, "invalid link")
		return
	}

	url, err := h.Services.GetURL(c, link)
	if err != nil {
		NewLinksError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if url == "" {
		NewLinksError(c, http.StatusInternalServerError, "no urls in result set")
		return
	}

	c.JSON(http.StatusOK, GetUrlResponse{
		Payload: &UrlPayload{Url: url},
	})
}
