package interfaces

import (
	"smart-blog/blog/articles/models"

	"github.com/gin-gonic/gin"
)

type ArticleRepositoryInterface interface {
	GetAll(c *gin.Context) ([]models.Article, error)                                           // Read all articles
	GetByID(c *gin.Context, id uint) (models.Article, error)                                   // Read an article by its ID
	Create(c *gin.Context, article models.Article) (models.Article, error)                     // Create a new article
	Update(c *gin.Context, id uint, article models.ArticleUpdateInput) (models.Article, error) // Update an existing article
	Delete(c *gin.Context, id uint) error
}
