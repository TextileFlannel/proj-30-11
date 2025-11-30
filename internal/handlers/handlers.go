package handlers

import (
	"net/http"
	"proj/internal/models"
	"proj/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Links(c *gin.Context) {
	var res models.LinksResponse
	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	req, err := h.service.Links(res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, req)
}

func (h *Handler) GetAllLinks(c *gin.Context) {
	req, err := h.service.GetAllLinks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, req)
}
