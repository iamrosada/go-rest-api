package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/go-rest-api/schemas"
)

// @BasePath /api/v1

// @Summary Create oportunity
// @Description Create a new job oportunity
// @Tags Oportunities
// @Accept json
// @Produce json
// @Param request body CreateOportunityRequest true "Request body"
// @Success 200 {object} CreateOportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /oportunity [post]
func CreateOportunityHandler(ctx *gin.Context) {
	request := CreateOportunityRequest{}

	ctx.BindJSON(&request)

	//logger.Infof("request received: %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validatione error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	oportunity := schemas.Oportunity{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&oportunity).Error; err != nil {
		logger.Errorf("error creating oportunity: %+v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating oportunity on database")
		return
	}
	sendSucess(ctx, "creating-oportunity", oportunity)
}
