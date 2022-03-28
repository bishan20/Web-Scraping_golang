package filter

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/store"
	"Web-Scraping_golang/utils/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateFilteredScrape(c *gin.Context) {
	var req models.CreateFilterRequest
	var combineData []models.FilterResponse

	var rsp models.CreateFilterResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Request"))
		return
	}
	layout := "2006-01-02"
	from, err := time.Parse(layout, req.FromDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse((err.Error())))
		return
	}

	to, err := time.Parse(layout, req.ToDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse((err.Error())))
		return
	}
	data, err := store.DBState.GetFilteredScrape(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Unable to create account"))
		return
	}
	for _, v := range data {
		item := models.FilterResponse{
			ID:           v.ID,
			UserID:       v.UserID,
			ScrappingUrl: v.ScrappingUrl,
			ScrappedData: v.ScrappedData,
			CreatedAt:    v.CreatedAt,
		}
		combineData = append(combineData, item)
	}
	rsp = models.CreateFilterResponse{
		Filter: combineData,
	}
	c.JSON(http.StatusOK, rsp)
}
