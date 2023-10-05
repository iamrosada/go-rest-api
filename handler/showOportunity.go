package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/schemas"
)

func ShowOportunityHandler(ctx *gin.Context) {

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest,
			errParamsIsRequired("id", "queryParameters").Error())
		return
	}

	oportunity := schemas.Oportunity{}
	if err := db.First(&oportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "oportunity not found")
		return
	}

	sendSucess(ctx, "show-oportunity", oportunity)
}
