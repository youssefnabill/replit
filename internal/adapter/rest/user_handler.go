
package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"hexagonal-backend/internal/app"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userApp *app.UserApp
}

// NewUserHandler creates a new user handler
func NewUserHandler(userApp *app.UserApp) *UserHandler {
	return &UserHandler{
		userApp: userApp,
	}
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userApp.CreateUser(c.Request.Context(), req.Email, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userApp.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ListUsers handles GET /users
func (h *UserHandler) ListUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	users, err := h.userApp.ListUsers(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
