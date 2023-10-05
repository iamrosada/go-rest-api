package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/iamrosada/go-rest-api/docs"
	"github.com/iamrosada/go-rest-api/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize Handler
	handler.InitializeHendler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	//Creating  a groupe of routes
	v1 := router.Group(basePath)
	{
		v1.GET("/oportunity", handler.ShowOportunityHandler)
		v1.PUT("/oportunity", handler.UpdateOportunityHandler)
		v1.POST("/oportunity", handler.CreateOportunityHandler)
		v1.DELETE("/oportunity", handler.DeleteOportunityHandler)
		v1.GET("/oportunities", handler.ListOportunitiesHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
