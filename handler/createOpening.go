package handler

import (
	"net/http"

	"github.com/alephjunio/go-opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("Error binding JSON: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := request.Validate()
	if err != nil {
		logger.Errorf("Validation error creating opening: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening :=
		schemas.Opening{
			Role:     request.Role,
			Company:  request.Company,
			Location: request.Location,
			Remote:   *request.Remote,
			Link:     request.Link,
			Salary:   request.Salary,
		}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	SendSucess(ctx, "Create Opening", opening)
}
