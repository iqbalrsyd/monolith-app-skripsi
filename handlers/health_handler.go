package handlers

import (
	"net/http"
	"time"

	"monolith-app/models"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check requests.
type HealthHandler struct{}

// NewHealthHandler creates a new HealthHandler instance.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health handles GET requests for health check.
func (h *HealthHandler) Health(c *gin.Context) {
	response := models.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Service:   "monolith-app",
		Version:   "1.0.0",
	}
	c.JSON(http.StatusOK, response)
}
