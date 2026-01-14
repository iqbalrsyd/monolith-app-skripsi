// Package models provides data structures for the monolith application.
package models

import "time"

// User represents a user in the system.
type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"created_at"`
}

// Product represents a product in the system.
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreateAt    time.Time `json:"created_at"`
}

// Order represents an order in the system.
type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	Products  []OrderItem `json:"products"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

// OrderItem represents an item in an order.
type OrderItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// HealthResponse represents the health check response.
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
	Version   string `json:"version"`
}
