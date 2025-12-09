package handler

import (
	"fmt"
	"net/http"

	"github.com/alephjunio/go-opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update Opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Param request body UpdateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening/{id} [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("Error binding JSON: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := request.Validate()
	if err != nil {
		logger.Errorf("Validation error opening: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "path").Error())
		return
	}

	opening := schemas.Opening{}

	err = db.First(&opening, id).Error

	if err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	err = db.Save(&opening).Error
	if err != nil {
		logger.Errorf("Error updating opening: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	SendSucess(ctx, "updating-opening", opening)

}
