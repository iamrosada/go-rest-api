package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOportunitiyHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
