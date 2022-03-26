package handler

import (
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
