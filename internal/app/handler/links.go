package handler

import (
	"github.com/BlackRRR/shortened-links/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ChangeURL(c *gin.Context) {
	var req ChangeUrlRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewLinksError(c, http.StatusBadRequest, "invalid req body")
		return
	}

	link, err := h.Links.ChangeURL(c, req.Url)
	if err != nil {
		models.NewLinksError(c, http.StatusInternalServerError, err.Error())
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

	url, err := h.Links.GetURL(c, link)
	if err != nil {
		models.NewLinksError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetUrlResponse{
		Payload: &UrlPayload{Url: url},
	})
}
