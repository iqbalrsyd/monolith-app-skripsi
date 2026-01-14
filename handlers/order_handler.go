package handlers

import (
	"monolith-app/models"
	"monolith-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderHandler handles order-related HTTP requests.
type OrderHandler struct {
	orderService *services.OrderService
}

// NewOrderHandler creates a new OrderHandler instance.
func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder handles POST requests to create a new order.
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var request struct {
		UserID string             `json:"user_id" binding:"required"`
		Items  []models.OrderItem `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.CreateOrder(request.UserID, request.Items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrders handles GET requests to retrieve orders by user ID.
func (h *OrderHandler) GetOrders(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id parameter is required"})
		return
	}

	orders, err := h.orderService.GetOrdersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
