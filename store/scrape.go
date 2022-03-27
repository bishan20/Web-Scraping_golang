package store

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/utils/channels"
)

func (state *State) CreateScrape(scrape models.Scrape) (models.Scrape, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = state.db.Debug().Model(&models.Scrape{}).Create(&scrape).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return scrape, nil
	}
	return models.Scrape{}, err
}

// get scrapped data details
func (state *State) GetScrape() ([]models.Scrape, error) {
	var err error
	var resp []models.Scrape
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = state.db.Model(&models.Scrape{}).Find(&resp).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return resp, nil
	}
	return []models.Scrape{}, err
}
