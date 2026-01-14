// Package services provides business logic for the monolith application.
package services

import (
	"fmt"
	"monolith-app/models"
	"time"
)

// UserService handles user-related business logic.
type UserService struct{}

// NewUserService creates a new UserService instance.
func NewUserService() *UserService {
	return &UserService{}
}

// GetAllUsers retrieves all users from the system.
func (s *UserService) GetAllUsers() ([]models.User, error) {
	users := []models.User{
		{
			ID:       "1",
			Name:     "John Doe",
			Email:    "john@example.com",
			CreateAt: time.Now(),
		},
		{
			ID:       "2",
			Name:     "Jane Smith",
			Email:    "jane@example.com",
			CreateAt: time.Now(),
		},
	}
	return users, nil
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	if id == "1" {
		return &models.User{
			ID:       "1",
			Name:     "John Doe",
			Email:    "john@example.com",
			CreateAt: time.Now(),
		}, nil
	}
	return nil, fmt.Errorf("user not found")
}
