package router

import (
	"net/http"

	"github.com/Cannedsans/GoCrud/auth"
	docs "github.com/Cannedsans/GoCrud/docs"
	"github.com/Cannedsans/GoCrud/handler"
	"github.com/gin-gonic/gin"
	saggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeAuthPath(router *gin.Engine, route string) {
	router.POST(route+"/auth", func(ctx *gin.Context) {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		ctx.BindJSON(&body)

		// just an fake validation
		if body.Email != "teste@teste.com" || body.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		token,_ := auth.GenerateJWT(1)

		ctx.JSON(http.StatusOK,gin.H{
			"token": token,
		})
	})
}

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	initializeAuthPath(router, basePath)
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1/")
	{
		v1.Use(auth.AuthRequired())

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

	}
	// initializa swagger

	router.GET("/swagger/*any", ginSwagger.WrapHandler(saggerfiles.Handler))
}
