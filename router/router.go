package router

import (
	"github.com/gin-gonic/gin"
)

func Inizialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.Run(":8080")

}
