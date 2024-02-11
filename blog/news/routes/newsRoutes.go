package routes

import (
	"smart-blog/blog/news/controllers"
	"smart-blog/blog/news/repositories/implementations"

	"github.com/gin-gonic/gin"
)

func RegisterNewsRoutes(router *gin.Engine) {

	newsRepo := implementations.NewNewsRepository()
	newsController := controllers.NewNewsController(newsRepo)

	router.GET("/news", newsController.GetAllNews)
	router.GET("/news/:id", newsController.GetNewsByID)
	router.POST("/news", newsController.CreateNews)
	router.PUT("/news/:id", newsController.UpdateNews)
	router.DELETE("/news/:id", newsController.DeleteNews)
}
