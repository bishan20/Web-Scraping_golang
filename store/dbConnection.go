package store

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/utils/token"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBState State

type State struct {
	db         *gorm.DB
	Config     *models.Config
	TokenMaker token.Maker
}

func ConnectionDb(config models.Config) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection problem:", err)
		panic("Cannot connect to database")
	}

	return DB

}

func Migration(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.Users{}, &models.Scrape{})
	if err != nil {
		log.Fatal("migration problem:", err)
	}

}

func Store(config models.Config, db *gorm.DB) {
	var err error
	Migration(db)
	tokenMaker, err := token.NewJWTMaker(config.AccessTokenSymmetricKey)
	if err != nil {
		fmt.Errorf("Cannot create token maker: %w", err)
		return
	}
	DBState = State{
		db:         db,
		Config:     &config,
		TokenMaker: tokenMaker,
	}

}
