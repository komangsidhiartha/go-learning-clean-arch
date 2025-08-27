package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sidhiartha.my.id/go-template/config"
)

func generateUrl(
	host string,
	port string,
	user string,
	password string,
	database string,
) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
}

func Connect(
	driver string,
	db config.DB,
) *gorm.DB {
	gormDb, err := gorm.Open(mysql.Open(generateUrl(db.Host, db.Port, db.User, db.Password, db.Database)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to database")
	return gormDb
}
