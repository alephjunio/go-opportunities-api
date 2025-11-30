package router

import (
	"github.com/gin-gonic/gin"
)

func Inizialize() {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"mensage": "pong",
		})
	})

}
