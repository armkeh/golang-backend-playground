package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Simple ping responder
func pingHandler(c *gin.Context) () {
	c.String(http.StatusOK, "pong")
}

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", pingHandler)

	return r
}

