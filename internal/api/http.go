package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/models"
	"github.com/michaelgbenle/rateApp/internal/ports"
)

type HTTPHandler struct {
	Repository ports.Repository
}

func NewHTTPHandler(repository ports.Repository) *HTTPHandler {
	return &HTTPHandler{
		Repository: repository,
	}
}

func (u *HTTPHandler) GetUserFromContext(c *gin.Context) (*models.User, error) {
	contextUser, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := contextUser.(*models.User)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}
