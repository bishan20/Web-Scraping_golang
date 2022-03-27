package store

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/utils/channels"
)

// Creating new user's account
func (state *State) CreateUser(user models.Users) (models.Users, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = state.db.Debug().Model(&models.Users{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.Users{}, err
}
