package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spargapees/users-api/service"
)

type Handler struct {
	services service.Service
}

func NewHandler(services service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	users := router.Group("/users")
	{
		users.POST("/")
		users.GET("/")
		users.PUT("/")

	}

	return router
}
