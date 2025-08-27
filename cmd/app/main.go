package main

import (
	"log"

	"sidhiartha.my.id/go-template/config"
	"sidhiartha.my.id/go-template/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("config.NewConfig error: %v", err)
	}

	app.Run(cfg)
}
