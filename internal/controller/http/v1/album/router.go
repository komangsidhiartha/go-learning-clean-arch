package album

import (
	"github.com/gin-gonic/gin"
	"sidhiartha.my.id/go-template/internal/usecase"
)

func NewRouter(router *gin.RouterGroup, uc *usecase.UseCase) {
	h := NewHandler(uc.AlbumUseCase)

	albumRouter := router.Group("/albums")
	{
		albumRouter.GET("", h.GetAlbums)
		albumRouter.GET("/:id", h.GetAlbumById)
	}
}
