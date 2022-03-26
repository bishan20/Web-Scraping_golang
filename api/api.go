package api

import (
	"Web-Scraping_golang/handler/hello"
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/repository"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config models.Config
	store  repository.Store
	//	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config models.Config, store repository.Store) (*Server, error) {
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot create token maker: %w", err)
	// }
	server := &Server{
		config: config,
		store:  store,
		//tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/hello", hello.ShowHello)
	server.router = router
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
