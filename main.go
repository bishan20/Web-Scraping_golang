package main

import (
	"Web-Scraping_golang/api"
	"Web-Scraping_golang/store"
	"Web-Scraping_golang/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Configuration could not be loaded:", err)
	}
	dbConn := store.ConnectionDb(config)
	store.Store(config, dbConn)
	api.Server(config)

	err = api.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error while starting server:", err)
	}
}
