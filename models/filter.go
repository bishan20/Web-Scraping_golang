package models

import "time"

type CreateFilterRequest struct {
	FromDate string `json:"from_date" `
	ToDate   string `json:"to_date" `
}

type CreateFilterResponse struct {
	Filter []FilterResponse `json:"scrapped"`
}

type FilterResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ScrappingUrl string    `json:"scrape_url"`
	ScrappedData string    `json:"scrapped_data"`
	CreatedAt    time.Time `json:"created_at"`
}
