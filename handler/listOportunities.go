package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/schemas"
)

// @BasePath /api/v1

// @Summary List Oportunities
// @Description List all job Oportunity
// @Tags Oportunities
// @Accept json
// @Produce json
// @Success 200 {object} ListOportunitiesResponse
// @Failure 500 {object} ErrorResponse
// @Router /Oportunities [get]
func ListOportunitiesHandler(ctx *gin.Context) {

	oportunities := []schemas.Oportunity{}

	if err := db.Find(&oportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing oportunities")
		return
	}

	sendSucess(ctx, "list-oportunities", oportunities)
}
