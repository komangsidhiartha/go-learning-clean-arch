package app

import (
	"gorm.io/gorm"
	"sidhiartha.my.id/go-template/config"
)

type Database struct {
	Gorm *gorm.DB
}

type App struct {
	DB     Database
	Config *config.Config
}
