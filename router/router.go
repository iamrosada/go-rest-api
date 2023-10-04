package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize the Route using default config
	router := gin.Default()

	initializeRoutes(router)

	// //Start running the API
	router.Run(":8080")
}
