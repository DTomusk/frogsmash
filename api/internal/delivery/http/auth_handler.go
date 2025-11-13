package http

import (
	"frogsmash/internal/container"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	RegisterUser(username, email, password string) error
	Login(username, password string) (string, error)
	RefreshToken(refreshToken string) (string, error)
}

type AuthHandler struct {
	AuthService AuthService
}

func NewAuthHandler(c *container.Container) *AuthHandler {
	return &AuthHandler{
		AuthService: c.AuthService,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Registers a new user with username, email, and password
// @Router       /register [post]
// @Accept       json
// @Produce      json
func (h *AuthHandler) Register(ctx *gin.Context) {
	// Implementation for user registration
}

// Login godoc
// @Summary      User login
// @Description  Logs in a user with username and password
// @Router       /login [post]
// @Accept       json
// @Produce      json
func (h *AuthHandler) Login(ctx *gin.Context) {
	// Implementation for user login
}

// RefreshToken godoc
// @Summary      Refresh JWT token
// @Description  Refreshes the JWT token using a refresh token
// @Router       /refresh [post]
// @Accept       json
// @Produce      json
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	// Implementation for refreshing JWT token
}
