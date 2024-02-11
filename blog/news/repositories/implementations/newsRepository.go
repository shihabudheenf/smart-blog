package implementations

import (
	"smart-blog/blog/news/models"
	dbutil "smart-blog/common"

	"github.com/gin-gonic/gin"
)

type NewsRepository struct{}

func NewNewsRepository() *NewsRepository {
	return &NewsRepository{}
}

func (repo *NewsRepository) GetAll(c *gin.Context) ([]models.News, error) {

	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return nil, err
	}

	var news []models.News

	result := gormDB.Find(&news)

	return news, result.Error
}

func (repo *NewsRepository) GetByID(c *gin.Context, id uint) (models.News, error) {

	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return models.News{}, err
	}

	var news models.News

	// Use `First` to find the news by ID. `First` adds a `WHERE` clause (e.g., `WHERE id = ?`)
	result := gormDB.First(&news, id)
	if result.Error != nil {
		return models.News{}, result.Error
	}

	return news, nil
}

func (repo *NewsRepository) Create(c *gin.Context, news models.News) (models.News, error) {

	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return models.News{}, err
	}

	// Create the news using the provided GORM DB connection
	if err := gormDB.Create(&news).Error; err != nil {
		return models.News{}, err
	}

	return news, nil
}

func (repo *NewsRepository) Update(c *gin.Context, id uint, updateData models.NewsUpdateInput) (models.News, error) {
	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return models.News{}, err
	}

	// Retrieve the existing news from the database
	var news models.News
	if err := gormDB.First(&news, id).Error; err != nil {
		return models.News{}, err
	}

	// Update the news's fields with the new values if they are provided (not nil)
	if updateData.Title != nil {
		news.Title = *updateData.Title
	}
	if updateData.Content != nil {
		news.Content = *updateData.Content
	}

	// Save the updated news back to the database
	if err := gormDB.Save(&news).Error; err != nil {
		return models.News{}, err
	}

	return news, nil
}

func (repo *NewsRepository) Delete(c *gin.Context, id uint) error {
	// Use the utility function to get the *gorm.DB from the Gin context
	gormDB, err := dbutil.GetDBFromContext(c)
	if err != nil {
		return err
	}

	// Perform the delete operation
	if err := gormDB.Delete(&models.News{}, id).Error; err != nil {
		return err
	}

	return nil
}
