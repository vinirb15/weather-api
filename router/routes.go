package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/vinirb15/go-api/docs"
	"github.com/vinirb15/go-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/weather", handler.ShowweatherHandler)
		v1.POST("/weather", handler.CreateweatherHandler)
		v1.DELETE("/weather", handler.DeleteweatherHandler)
		v1.PUT("/weather", handler.UpdateweatherHandler)
		v1.GET("/weathers", handler.ListweathersHandler)

	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
