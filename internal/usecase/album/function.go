package album

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"sidhiartha.my.id/go-template/internal/entity/dto"
)

func (u albumUseCase) GetAlbums(ctx *gin.Context) ([]dto.Album, error) {
	albums, err := u.Repository.GetAlbums(ctx)

	if err != nil {
		return []dto.Album{}, fmt.Errorf("failed to get albums: %w", err)
	}

	return albums, nil
}

func (u albumUseCase) GetAlbumById(ctx *gin.Context, id string) (dto.Album, error) {
	album, err := u.Repository.GetAlbumByID(ctx, id)

	if err != nil {
		return dto.Album{}, fmt.Errorf("failed to get album: %w", err)
	}

	return album, nil
}
