package handler

import (
	"github.com/BlackRRR/shortened-Links/internal/app/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Links *services.LinksService
}

func NewHandler(service *services.Services) *Handler {
	return &Handler{Links: service.Links}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/", h.ChangeURL)
	router.GET("/:link", h.GetURL)

	return router
}
