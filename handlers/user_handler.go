// Package handlers provides HTTP request handlers for the monolith application.
package handlers

import (
	"monolith-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers handles GET requests to retrieve all users.
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser handles GET requests to retrieve a user by ID.
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
