package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/schemas"
)

// @BasePath /api/v1

// @Summary Update Oportunity
// @Description Update a job Oportunity
// @Tags Oportunities
// @Accept json
// @Produce json
// @Param id query string true "Oportunity Identification"
// @Param Oportunity body UpdateOportunityRequest true "Oportunity data to Update"
// @Success 200 {object} UpdateOportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /oportunity [put]
func UpdateOportunityHandler(ctx *gin.Context) {

	request := UpdateOportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return

	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest,
			errParamsIsRequired("id", "queryParameters").Error())
		return
	}

	oportunity := schemas.Oportunity{}
	if err := db.First(&oportunity, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest,
			"oportunity not found")

		return
	}

	if request.Role != "" {
		oportunity.Role = request.Role
	}

	if request.Company != "" {
		oportunity.Company = request.Company
	}

	if request.Location != "" {
		oportunity.Location = request.Location
	}

	if request.Link != "" {
		oportunity.Link = request.Link
	}

	if request.Remote != nil {
		oportunity.Remote = *request.Remote
	}

	if request.Salary <= 0 {
		oportunity.Salary = request.Salary
	}

	//save oportunity

	if err := db.Save(&oportunity).Error; err != nil {
		logger.Errorf("error updating oportunity: %+v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updatings oportunity on database")
		return
	}
	sendSucess(ctx, "updating-oportunity", oportunity)

}
