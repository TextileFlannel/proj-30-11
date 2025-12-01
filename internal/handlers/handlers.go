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
	var req models.LinksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Links(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllLinks(c *gin.Context) {
	res, err := h.service.GetAllLinks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) ReportLinks(c *gin.Context) {
	var req models.ReportLinksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buf, err := h.service.ReportLinks(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.DataFromReader(http.StatusOK,
		int64(buf.Len()),
		"application/pdf",
		&buf,
		map[string]string{"Content-Disposition": `attachment; filename="generated.pdf"`},
	)
}
