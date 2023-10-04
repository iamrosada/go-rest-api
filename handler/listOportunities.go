package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOportunitiesHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
