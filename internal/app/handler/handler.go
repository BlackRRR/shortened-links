package handler

import (
	"github.com/BlackRRR/shortened-links/internal/app/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*services.Services
}

func NewHandler(service *services.Services) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/", h.ChangeURL)
	router.GET("/:link", h.GetURL)

	return router
}
