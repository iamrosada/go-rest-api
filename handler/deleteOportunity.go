package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/schemas"
)

// @BasePath /api/v1

// @Summary Delete oportunity
// @Description Delete a new job oportunity
// @Tags Oportunities
// @Accept json
// @Produce json
// @Param id query string  true "Oportunity identification"
// @Success 200 {object} DeleteOportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /oportunity [delete]
func DeleteOportunityHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest,
			errParamsIsRequired("id", "queryParameters").Error())
		return
	}

	oportunity := schemas.Oportunity{}

	//Find the Oportunity by id

	if err := db.First(&oportunity, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest,
			fmt.Sprintf("Oportunity with id: %s not found", id))

		return
	}

	//Delete Oportunity
	if err := db.Delete(&oportunity, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest,
			fmt.Sprintf("Error deleting oportunity with id: %s", id))

		return
	}

	sendSucess(ctx, "delete-oportunity", oportunity)
}
