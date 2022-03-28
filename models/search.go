package models

import "time"

type SearchResponse struct {
	SearchList []SearchDataResponse `json:"scrapped"`
}

type SearchDataResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ScrappingUrl string    `json:"scrape_url"`
	ScrappedData string    `json:"scrapped_data"`
	CreatedAt    time.Time `json:"created_at"`
}
