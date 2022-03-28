package search

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/store"
	"Web-Scraping_golang/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	var rsp models.SearchResponse
	var combine []models.SearchDataResponse
	q := c.Request.URL.Query().Get("q")
	data, err := store.DBState.GetSearchedScrape(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("error while searching"))
		return
	}
	for _, v := range data {
		items := models.SearchDataResponse{
			ID:           v.ID,
			UserID:       v.UserID,
			ScrappingUrl: v.ScrappingUrl,
			ScrappedData: v.ScrappedData,
			CreatedAt:    v.CreatedAt,
		}
		combine = append(combine, items)
	}
	rsp = models.SearchResponse{
		SearchList: combine,
	}
	c.JSON(http.StatusOK, rsp)
}
