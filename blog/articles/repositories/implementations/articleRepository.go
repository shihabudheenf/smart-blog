package implementations

import (
	"smart-blog/blog/articles/models"
	dbutil "smart-blog/common"

	"github.com/gin-gonic/gin"
)

type ArticleRepository struct{}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (repo *ArticleRepository) GetAll(c *gin.Context) ([]models.Article, error) {

	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	result := gormDB.Find(&articles)

	return articles, result.Error
}

func (repo *ArticleRepository) GetByID(c *gin.Context, id uint) (models.Article, error) {

	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return models.Article{}, err
	}

	var article models.Article

	// Use `First` to find the article by ID. `First` adds a `WHERE` clause (e.g., `WHERE id = ?`)
	result := gormDB.Preload("Category").First(&article, id)
	if result.Error != nil {
		return models.Article{}, result.Error
	}

	return article, nil
}

func (repo *ArticleRepository) Create(c *gin.Context, article models.Article) (models.Article, error) {

	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return models.Article{}, err
	}

	// Create the article using the provided GORM DB connection
	if err := gormDB.Create(&article).Error; err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func (repo *ArticleRepository) Update(c *gin.Context, id uint, updateData models.ArticleUpdateInput) (models.Article, error) {
	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return models.Article{}, err
	}

	// Retrieve the existing article from the database
	var article models.Article
	if err := gormDB.First(&article, id).Error; err != nil {
		return models.Article{}, err
	}

	// Update the article's fields with the new values if they are provided (not nil)
	if updateData.Title != nil {
		article.Title = *updateData.Title
	}
	if updateData.CategoryID != nil {
		article.CategoryID = *updateData.CategoryID
	}
	if updateData.Author != nil {
		article.Author = *updateData.Author
	}

	// Save the updated article back to the database
	if err := gormDB.Save(&article).Error; err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func (repo *ArticleRepository) Delete(c *gin.Context, id uint) error {
	// Use the utility function to get the *gorm.DB from the Gin context
	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return err
	}

	// Perform the delete operation
	if err := gormDB.Delete(&models.Article{}, id).Error; err != nil {
		return err
	}

	return nil
}
