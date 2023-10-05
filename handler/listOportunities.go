package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/schemas"
)

func ListOportunitiesHandler(ctx *gin.Context) {

	oportunities := []schemas.Oportunity{}

	if err := db.Find(&oportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing oportunities")
		return
	}

	sendSucess(ctx, "list-oportunities", oportunities)
}
