package handler

import (
	"jas/src/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
// Create Ping function in handler.go
// Ping godoc
// @Summary Ping the server
// @Description ping the server to check if it's alive
// @Tags ping
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /ping [get]
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
