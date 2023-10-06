package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {

	ctx.Header("Content-type", "application")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSucess(ctx *gin.Context, option string, data interface{}) {

	ctx.Header("Content-type", "application")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessefull", option),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOportunityResponse struct {
	Message string                     `json:"message"`
	Data    schemas.OportunityResponse `json:"data"`
}

type DeleteOportunityResponse struct {
	Message string                     `json:"message"`
	Data    schemas.OportunityResponse `json:"data"`
}

type ShowOportunityResponse struct {
	Message string                     `json:"message"`
	Data    schemas.OportunityResponse `json:"data"`
}

type ListOportunitiesResponse struct {
	Message string                       `json:"message"`
	Data    []schemas.OportunityResponse `json:"data"`
}
type UpdateOportunityResponse struct {
	Message string                     `json:"message"`
	Data    schemas.OportunityResponse `json:"data"`
}
