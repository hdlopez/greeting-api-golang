package api

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Api struct {
}

func New() *Api {
	return new(Api)
}

func (api *Api) Run() {
	r := gin.Default()

	r.GET("/greeting", cors(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"owner":      "Horacio Lopez",
			"greeting":   "IT Expert / Technical Manager @MercadoLibre",
			"repository": "https://github.com/hdlopez/greeting-api-golang",
		})
	})

	r.Run(api.port()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (api *Api) port() string {
	// heroku variable
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}
