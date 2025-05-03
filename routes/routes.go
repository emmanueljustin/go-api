package routes

import (
	"github.com/emmanueljustin/go-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/sample", controllers.GetSampleData)
}
