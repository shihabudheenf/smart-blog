package routes

import (
	"smart-blog/blog/articles/controllers"
	"smart-blog/blog/articles/repositories/implementations"

	"github.com/gin-gonic/gin"
)

func RegisterArticleRoutes(router *gin.Engine) {

	articleRepo := implementations.NewArticleRepository()
	articleController := controllers.NewArticleController(articleRepo)

	router.GET("/articles", articleController.GetAllArticles)
	router.GET("/articles/:id", articleController.GetArticleByID)
	router.POST("/articles", articleController.CreateArticle)
	router.PUT("/articles/:id", articleController.UpdateArticle)
	router.DELETE("/articles/:id", articleController.DeleteArticle)
}
