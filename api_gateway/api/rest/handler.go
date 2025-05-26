package rest

import (
	"api_gateway/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	service *service.Service
	logger  *zap.Logger
}

func RegisterRoutes(r *gin.Engine, svc *service.Service, logger *zap.Logger) {
	h := &Handler{svc, logger}
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", h.Ping)
	}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
