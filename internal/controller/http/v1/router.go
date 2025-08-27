package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sidhiartha.my.id/go-template/internal/controller/http/v1/album"

	"sidhiartha.my.id/go-template/internal/usecase"
)

func NewRouter(handler *gin.Engine, useCase *usecase.UseCase) {
	h := handler.Group("/api/v1")
	{
		h.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })

		album.NewRouter(h, useCase)
	}
}
