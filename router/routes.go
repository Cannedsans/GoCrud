package router

import (
	"github.com/Cannedsans/GoCrud/handler"
	"github.com/gin-gonic/gin"
	docs "github.com/Cannedsans/GoCrud/docs"
	saggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

	}
	// initializa swagger

	router.GET("/swagger/*any",ginSwagger.WrapHandler(saggerfiles.Handler))
}
