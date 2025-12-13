package services

import (
	"fmt"
	"time"
	"monolith-app/models"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

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