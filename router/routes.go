package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1/")
	{
		v1.GET("openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensage": "GET - List opening",
			})
		})
		v1.POST("opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensage": "POST - Create opening",
			})
		})
		v1.GET("opening/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensage": "SHOW - Show opening",
			})
		})
		v1.PUT("opening/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensage": "PUT - Update opening",
			})
		})
		v1.DELETE("opening/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensage": "DELETE - Delete opening",
			})
		})

	}

}
