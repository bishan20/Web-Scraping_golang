package scrape

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/store"
	scrapping "Web-Scraping_golang/utils/Scrapping"
	"Web-Scraping_golang/utils/response"
	"Web-Scraping_golang/utils/token"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	authorizationPayloadKey = "authorization_payload"
)

func CreateScrape(c *gin.Context) {
	var req models.CreateScrapeRequest
	var rsp models.CreateScrapeResponse
	var created []models.ScrapeResponse
	var wg sync.WaitGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Request"))
		return
	}
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	wg.Add(len(req.ScrappingUrl))
	d := make(chan models.Scrape)
	for _, v := range req.ScrappingUrl {
		go scrapping.ScrappingUrl(v, authPayload.UserId, d, &wg)
		data, err := store.DBState.CreateScrape(<-d)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("error while scrapping"))
			return
		}
		item := models.ScrapeResponse{
			ID:           data.ID,
			UserID:       data.UserID,
			ScrappingUrl: data.ScrappingUrl,
			ScrappedData: data.ScrappedData,
			CreatedAt:    data.CreatedAt,
		}
		created = append(created, item)

	}
	rsp = models.CreateScrapeResponse{
		Scrapped: created,
	}
	wg.Wait()

	c.JSON(http.StatusCreated, rsp)
}

func GetScrape(c *gin.Context) {

	var rsp models.CreateScrapeResponse
	var created []models.ScrapeResponse
	data, err := store.DBState.GetScrape()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("error while scrapping"))
		return
	}
	for _, v := range data {

		item := models.ScrapeResponse{
			ID:           v.ID,
			UserID:       v.UserID,
			ScrappingUrl: v.ScrappingUrl,
			ScrappedData: v.ScrappedData,
			CreatedAt:    v.CreatedAt,
		}
		created = append(created, item)

	}
	rsp = models.CreateScrapeResponse{
		Scrapped: created,
	}

	c.JSON(http.StatusOK, rsp)
}
