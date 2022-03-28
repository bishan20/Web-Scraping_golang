package users

import (
	"Web-Scraping_golang/models"
	"Web-Scraping_golang/store"

	"Web-Scraping_golang/utils/password"
	"Web-Scraping_golang/utils/response"
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
)

func newUser(u models.Users) models.UsersResponse {
	return models.UsersResponse{
		ID:        u.ID,
		UserName:  u.UserName,
		CreatedAt: u.CreatedAt,
	}
}
func UserLogin(c *gin.Context) {
	var req models.LoginRequest
	var rep models.LoginResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Request"))
		return
	}

	user, err := store.DBState.GetUser(req.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, response.ErrorResponse("User not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	err = password.CheckPassword(req.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse("Password Incorrect"))
		return
	}
	token, err := store.DBState.TokenMaker.CreateToken(user.UserName, user.ID, store.DBState.Config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(err.Error()))
		return
	}
	rep = models.LoginResponse{
		Token: token,
		User:  newUser(user),
	}

	c.JSON(http.StatusOK, rep)
}
