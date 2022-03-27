package users

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/store"
	"Web-Scraping_golang/utils/password"
	"Web-Scraping_golang/utils/response"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var req models.CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Request"))
		return
	}
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Unable to Hash"))
		return
	}
	user := models.Users{
		UserName: req.UserName,
		Password: hashedPassword,
	}
	users, err := store.DBState.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Unable to create account"))
		return
	}
	c.JSON(http.StatusCreated, users)
}
