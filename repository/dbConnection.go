package repository

import (
	"Web-Scraping_golang/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDb(config models.Config) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	return DB

}
