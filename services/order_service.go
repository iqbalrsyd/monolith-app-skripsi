package services

import (
	"fmt"
	"math"
	"monolith-app/models"
	"time"
)

// OrderService handles order-related business logic.
type OrderService struct {
	userService    *UserService
	productService *ProductService
}

// NewOrderService creates a new OrderService instance.
func NewOrderService(userService *UserService, productService *ProductService) *OrderService {
	return &OrderService{
		userService:    userService,
		productService: productService,
	}
}

// CreateOrder creates a new order for a user.
func (s *OrderService) CreateOrder(userID string, items []models.OrderItem) (*models.Order, error) {
	total := 0.0
	for _, item := range items {
		product, err := s.productService.GetProductByID(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product %s not found", item.ProductID)
		}
		total += float64(item.Quantity) * product.Price
	}

	total = math.Round(total*100) / 100

	order := &models.Order{
		ID:        "ORD-" + fmt.Sprintf("%d", time.Now().Unix()),
		UserID:    userID,
		Products:  items,
		Total:     total,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	return order, nil
}

// GetOrdersByUserID retrieves all orders for a specific user.
func (s *OrderService) GetOrdersByUserID(userID string) ([]models.Order, error) {
	orders := []models.Order{
		{
			ID:     "ORD-1",
			UserID: userID,
			Products: []models.OrderItem{
				{ProductID: "1", Quantity: 1, Price: 999.99},
			},
			Total:     999.99,
			Status:    "completed",
			CreatedAt: time.Now().Add(-24 * time.Hour),
		},
	}
	return orders, nil
}
