package hello

import (
	"github.com/gin-gonic/gin"
)

func ShowHello(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hello",
	})
}
