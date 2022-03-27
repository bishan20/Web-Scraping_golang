package api

import (
	"Web-Scraping_golang/handler/hello"
	"Web-Scraping_golang/handler/users"
	"Web-Scraping_golang/models"

	"github.com/gin-gonic/gin"
)

var routers *gin.Engine

func Server(config models.Config) {

	setupRouter()

}

func setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/hello", hello.ShowHello)
	router.POST("/create-user", users.CreateAccount)
	router.POST("/login", users.UserLogin)
	routers = router

}
func Start(address string) error {
	return routers.Run(address)
}
