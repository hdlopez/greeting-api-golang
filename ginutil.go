package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func addCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Authorization, X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Allow-Request-Method")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	c.Header("Allow", "GET, POST, OPTIONS, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func port() string {
	// heroku variable
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}
