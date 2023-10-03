package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize the Route using default config
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	//Start running the API
	router.Run()
}
