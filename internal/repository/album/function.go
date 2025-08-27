package album

import (
	"context"

	"sidhiartha.my.id/go-template/internal/entity/dto"
)

func (a albumRepository) GetAlbums(ctx context.Context) ([]dto.Album, error) {
	var albums []dto.Album
	err := a.DB.WithContext(ctx).Model(&albums).Find(&albums).Error

	return albums, err
}

func (a albumRepository) GetAlbumByID(ctx context.Context, id string) (dto.Album, error) {
	var album dto.Album
	err := a.DB.WithContext(ctx).Model(&album).Where("id = ?", id).Find(&album).Error

	return album, err
}
