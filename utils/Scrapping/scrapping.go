package scrapping

import (
	"Web-Scraping_golang/models"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func ScrappingUrl(url string, userId uint, d chan models.Scrape, wg *sync.WaitGroup) {
	defer wg.Done()
	c := colly.NewCollector()
	title := ""

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.Visit(url)
	now := time.Now()

	d <- models.Scrape{
		UserID:       userId,
		ScrappingUrl: url,
		ScrappedData: title,
		CreatedAt:    now,
	}

}
