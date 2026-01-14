package services

import (
	"fmt"
	"monolith-app/models"
	"time"
)

// ProductService handles product-related business logic.
type ProductService struct{}

// NewProductService creates a new ProductService instance.
func NewProductService() *ProductService {
	return &ProductService{}
}

// GetAllProducts retrieves all products from the system.
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	products := []models.Product{
		{
			ID:          "1",
			Name:        "Laptop",
			Description: "High-performance laptop",
			Price:       999.99,
			Stock:       10,
			CreateAt:    time.Now(),
		},
		{
			ID:          "2",
			Name:        "Mouse",
			Description: "Wireless optical mouse",
			Price:       29.99,
			Stock:       50,
			CreateAt:    time.Now(),
		},
	}
	return products, nil
}

// GetProductByID retrieves a product by its ID.
func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	if id == "1" {
		return &models.Product{
			ID:          "1",
			Name:        "Laptop",
			Description: "High-performance laptop",
			Price:       999.99,
			Stock:       10,
			CreateAt:    time.Now(),
		}, nil
	}
	return nil, fmt.Errorf("product not found")
}
