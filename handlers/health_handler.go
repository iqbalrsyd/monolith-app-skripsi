package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"monolith-app/models"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	response := models.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Service:   "monolith-app",
		Version:   "1.0.0",
	}
	c.JSON(http.StatusOK, response)
}