package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3333")
}