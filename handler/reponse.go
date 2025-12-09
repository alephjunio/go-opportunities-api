package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, menssage string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": menssage,
		"code":    code,
	})
}

func SendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s success", op),
		"data":    data,
		"code":    http.StatusOK,
	})
}
