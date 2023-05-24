package handler

import (
	"booking_service_api/models"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService  models.UserService
	TokenService models.TokenService
}

type Config struct {
	R            *gin.Engine
	UserService  models.UserService
	TokenService models.TokenService
	BaseURL      string
}

func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	g := c.R.Group("/api")
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/tokens", h.Tokens)
}
