package store

import (
	"Web-Scraping_golang/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBState State

type State struct {
	db     *gorm.DB
	Config *models.Config
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
	Migration(db)
	DBState = State{
		db:     db,
		Config: &config,
	}

}
