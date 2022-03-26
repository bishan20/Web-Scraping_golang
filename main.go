package main

import (
	"Web-Scraping_golang/api"
	"Web-Scraping_golang/repository"
	"Web-Scraping_golang/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Configuration could not be loaded:", err)
	}
	conn := repository.ConnectionDb(config)
	store := repository.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
