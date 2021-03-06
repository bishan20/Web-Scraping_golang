package models

import "time"

type CreateScrapeRequest struct {
	ScrappingUrl []string `json:"scrape_url"`
}

type CreateScrapeResponse struct {
	Scrapped []ScrapeResponse `json:"scrapped"`
}

type ScrapeResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ScrappingUrl string    `json:"scrape_url"`
	ScrappedData string    `json:"scrapped_data"`
	CreatedAt    time.Time `json:"created_at"`
}
