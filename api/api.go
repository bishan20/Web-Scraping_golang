package api

import (
	"Web-Scraping_golang/handler/filter"
	"Web-Scraping_golang/handler/hello"
	"Web-Scraping_golang/handler/scrape"
	"Web-Scraping_golang/handler/search"
	"Web-Scraping_golang/handler/users"
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/store"

	"github.com/gin-gonic/gin"
)

var routers *gin.Engine

func Server(config models.Config) {

	setupRouter()

}

func setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	scrapeRoute := router.Group("/").Use(authMiddleware(store.DBState.TokenMaker))
	router.GET("/hello", hello.ShowHello)
	router.POST("/create-user", users.CreateAccount)
	router.POST("/login", users.UserLogin)
	scrapeRoute.POST("/scrape", scrape.CreateScrape)
	scrapeRoute.GET("/scrape", scrape.GetScrape)
	router.POST("/filter", filter.CreateFilteredScrape)
	router.GET("/search", search.Search)
	routers = router

}
func Start(address string) error {
	return routers.Run(address)
}
