package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOportunitiyHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST opening",
	})
}
func DeleteOportunitiyHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}
func ShowOportunitiyHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func ListOportunitiesHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
func UpdateOportunitiyHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}
