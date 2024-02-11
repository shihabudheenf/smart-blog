package interfaces

import (
	"smart-blog/blog/news/models"

	"github.com/gin-gonic/gin"
)

type NewsRepositoryInterface interface {
	GetAll(c *gin.Context) ([]models.News, error)                                           // Read all newss
	GetByID(c *gin.Context, id uint) (models.News, error)                                   // Read an news by its ID
	Create(c *gin.Context, news models.News) (models.News, error)                     // Create a new news
	Update(c *gin.Context, id uint, news models.NewsUpdateInput) (models.News, error) // Update an existing news
	Delete(c *gin.Context, id uint) error
}
