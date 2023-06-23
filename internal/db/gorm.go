package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/yosa12978/mockPizza/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		log.Println(dsn)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		migrate()
	})
	return db
}

func migrate() {
	db.AutoMigrate(&domain.Pizza{})
	db.AutoMigrate(&domain.Usr{})
	db.AutoMigrate(&domain.Order{})
}
