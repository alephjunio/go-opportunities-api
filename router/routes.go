package router

import (
	docs "github.com/alephjunio/go-opportunities-api/docs"
	"github.com/alephjunio/go-opportunities-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	basePath := "/api/v1/"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("openings", handler.ListOpeningsHandler)
		v1.POST("opening", handler.CreateOpeningHandler)
		v1.GET("opening/:id", handler.ShowOpeningHandler)
		v1.PUT("opening/:id", handler.UpdateOpeningHandler)
		v1.DELETE("opening/:id", handler.DeleteOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
