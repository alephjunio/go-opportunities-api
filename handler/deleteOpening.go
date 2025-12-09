package handler

import (
	"fmt"
	"net/http"

	"github.com/alephjunio/go-opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "path").Error())
		return
	}

	opening := schemas.Opening{}

	err := db.First(&opening, id).Error
	if err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	db.Delete(&opening)
	if err != nil {
		SendError(ctx, http.StatusBadRequest, fmt.Sprintf("error deleting opening with id: %v", id))
		return
	}

	SendSucess(ctx, "deleted-opening", opening)

}
