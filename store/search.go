package store

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/utils/channels"
)

func (state *State) GetSearchedScrape(url string) ([]models.Scrape, error) {
	var err error
	var resp []models.Scrape
	done := make(chan bool)
	go func(ch chan<- bool) {

		err = state.db.Model(&models.Scrape{}).Where("scrapping_url = ?", url).Find(&resp).Error
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
