package handler

import (
	"fmt"
	"net/http"

	"github.com/alephjunio/go-opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
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
	SendSucess(ctx, "show-opening", opening)
}
