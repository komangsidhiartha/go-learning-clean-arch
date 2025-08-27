package album

import (
	"sidhiartha.my.id/go-template/config"
	"sidhiartha.my.id/go-template/internal/repository/album"
)

func NewAlbumUseCase(cfg *config.Config, repository album.Repository) UseCase {
	return &albumUseCase{
		Cfg:        cfg,
		Repository: repository,
	}
}
