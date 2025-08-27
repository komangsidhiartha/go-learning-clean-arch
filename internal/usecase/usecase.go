package usecase

import (
	album2 "sidhiartha.my.id/go-template/internal/repository/album"
	"sidhiartha.my.id/go-template/internal/usecase/album"
	"sidhiartha.my.id/go-template/pkg/app"
)

type UseCase struct {
	App          *app.App
	AlbumUseCase album.UseCase
}

func NewUseCase(apps *app.App) *UseCase {
	albumRepository := album2.NewRepository(apps.DB.Gorm)

	albumUseCase := album.NewAlbumUseCase(apps.Config, albumRepository)

	useCase := &UseCase{
		App:          apps,
		AlbumUseCase: albumUseCase,
	}

	return useCase
}
