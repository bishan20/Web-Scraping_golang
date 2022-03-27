package api

import (
	"Web-Scraping_golang/handler/hello"
	"Web-Scraping_golang/handler/users"

	"github.com/gin-gonic/gin"
)

var routers *gin.Engine

func Server() {
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot create token maker: %w", err)
	// }
	setupRouter()
	return
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
