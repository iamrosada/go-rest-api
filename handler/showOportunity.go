package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/schemas"
)

// @BasePath /api/v1

// @Summary Show oportunity
// @Description Show a job oportunity
// @Tags Oportunities
// @Accept json
// @Produce json
// @Param id query string  true "Oportunity identification"
// @Success 200 {object} ShowOportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /oportunity [get]
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
