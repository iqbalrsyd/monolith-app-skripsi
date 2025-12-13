package services

import (
	"fmt"
	"time"
	"monolith-app/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

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