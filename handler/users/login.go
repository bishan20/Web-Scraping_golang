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

func UserLogin(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Request"))
		return
	}

	user, err := store.DBState.GetUser(req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, response.ErrorResponse("User not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	err = password.CheckPassword(req.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}
	// accessToken, err := server.tokenMaker.CreateToken(
	// 	user.Username,
	// 	user.ID,
	// 	server.config.AccessTokenDuration,
	// )
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }
	// res = loginRes{
	// 	User:  newUserResponse(user),
	// 	Token: accessToken,
	// }
	// c.JSON(http.StatusOK, res)
}
