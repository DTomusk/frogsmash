package http

import (
	"database/sql"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/dto"

	"context"

	"github.com/gin-gonic/gin"
)

// TODO: consider how to make this testable, *sql.DB is a concrete type
type AuthService interface {
	RegisterUser(username, email, password string, ctx context.Context, db repos.DBTX) error
	Login(username, password string, ctx context.Context, db repos.DBWithTxStarter) (string, string, error)
	RefreshToken(refreshToken string, ctx context.Context, db repos.DBTX) (string, error)
}

type AuthHandler struct {
	AuthService AuthService
	db          *sql.DB
}

func NewAuthHandler(c *container.Container) *AuthHandler {
	return &AuthHandler{
		AuthService: c.AuthService,
		db:          c.DB,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Registers a new user with username, email, and password
// @Router       /register [post]
// @Accept       json
// @Produce      json
// @Param        user  body  dto.UserRegistrationRequest  true  "User registration payload"
func (h *AuthHandler) Register(ctx *gin.Context) {
	var req dto.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.AuthService.RegisterUser(req.Username, req.Email, req.Password, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary      User login
// @Description  Logs in a user with username and password
// @Router       /login [post]
// @Accept       json
// @Produce      json
// @Param        user  body  dto.UserLoginRequest  true  "User login payload"
// @Success      200   {object}  dto.UserLoginResponse
func (h *AuthHandler) Login(ctx *gin.Context) {
	// Implementation for user login
	var req dto.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	jwt, refreshToken, err := h.AuthService.Login(req.Username, req.Password, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res := dto.UserLoginResponse{
		JWT:          jwt,
		RefreshToken: refreshToken,
	}
	ctx.JSON(200, res)
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
