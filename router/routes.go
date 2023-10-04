package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/handler"
)

func initializeRoutes(router *gin.Engine) {

	//Creating  a groupe of routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOportunitiyHandler)
		v1.PUT("/opening", handler.UpdateOportunitiyHandler)
		v1.POST("/opening", handler.CreateOportunitiyHandler)
		v1.DELETE("/opening", handler.DeleteOportunitiyHandler)
		v1.GET("/opening/o", handler.ListOportunitiesHandler)
	}
}
