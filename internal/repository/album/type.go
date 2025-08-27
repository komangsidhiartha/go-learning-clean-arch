package album

import (
	"context"

	"gorm.io/gorm"
	"sidhiartha.my.id/go-template/internal/entity/dto"
)

type Repository interface {
	GetAlbums(ctx context.Context) ([]dto.Album, error)
	GetAlbumByID(ctx context.Context, id string) (dto.Album, error)
}

type albumRepository struct {
	DB *gorm.DB
}
