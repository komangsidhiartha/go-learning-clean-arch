package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sidhiartha.my.id/go-template/internal/usecase/album"
)

type AlbumHandler interface {
	GetAlbums(ctx *gin.Context)
	GetAlbumById(ctx *gin.Context)
}

type albumHandler struct {
	albumUseCase album.UseCase
}

func NewHandler(albumUseCase album.UseCase) AlbumHandler {
	return &albumHandler{
		albumUseCase: albumUseCase,
	}
}

func (a albumHandler) GetAlbums(ctx *gin.Context) {
	albums, err := a.albumUseCase.GetAlbums(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"albums": albums})
}

func (a albumHandler) GetAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")

	album, err := a.albumUseCase.GetAlbumById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"album": album})
}
