package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"sidhiartha.my.id/go-template/config"
	v1 "sidhiartha.my.id/go-template/internal/controller/http/v1"
	"sidhiartha.my.id/go-template/internal/usecase"
	"sidhiartha.my.id/go-template/pkg/app"
	"sidhiartha.my.id/go-template/pkg/database"
)

func Run(cfg *config.Config) {
	fmt.Printf("configs: %+v\n", cfg)

	apps := app.App{
		DB: app.Database{
			Gorm: database.Connect(
				"mysql",
				cfg.DB,
			),
		},
		Config: cfg,
	}

	fmt.Printf("app: %+v\n", apps)

	useCase := usecase.NewUseCase(&apps)

	fmt.Printf("usecase: %+v\n", useCase)

	handler := gin.New()

	v1.NewRouter(handler, useCase)

	handler.Run("localhost:8081")
}
