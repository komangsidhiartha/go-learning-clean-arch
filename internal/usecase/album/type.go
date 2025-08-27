package album

import (
	"github.com/gin-gonic/gin"
	"sidhiartha.my.id/go-template/config"
	"sidhiartha.my.id/go-template/internal/entity/dto"
	"sidhiartha.my.id/go-template/internal/repository/album"
)

type UseCase interface {
	GetAlbums(ctx *gin.Context) ([]dto.Album, error)
	GetAlbumById(ctx *gin.Context, id string) (dto.Album, error)
}

type albumUseCase struct {
	Cfg        *config.Config
	Repository album.Repository
}
