package controllers

import (
	"net/http"
	"smart-blog/blog/news/models"
	"smart-blog/blog/news/repositories/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewsController struct {
	repo interfaces.NewsRepositoryInterface
}

func NewNewsController(repo interfaces.NewsRepositoryInterface) *NewsController {
	return &NewsController{
		repo: repo,
	}
}

func (ac *NewsController) GetAllNews(c *gin.Context) {
	news, err := ac.repo.GetAll(c) // Pass the Gin context here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, news)
}

func (ac *NewsController) GetNewsByID(c *gin.Context) {

	// Extract the 'id' path parameter and convert it to an unsigned integer
	idStr := c.Param("id")                      // 'id' is the name of the path parameter
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID format"})
		return
	}

	news, err := ac.repo.GetByID(c, uint(id)) // Note the conversion of id to uint
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, news)
}

// CreateNews handles the POST request for creating a new news
func (ac *NewsController) CreateNews(c *gin.Context) {
	var news models.News

	// Bind the JSON payload to the News struct
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the repository's Create method, passing the bound news
	_, err := ac.repo.Create(c, news)
	if err != nil {
		// Handle errors that may occur during news creation
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If successful, respond with the created news
	c.Status(http.StatusCreated)

}

func (ac *NewsController) UpdateNews(c *gin.Context) {
	// Assuming "id" is a parameter in the route
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid news ID"})
		return
	}

	// Bind the request body to the News struct
	var updatedNews models.NewsUpdateInput
	if err := c.ShouldBindJSON(&updatedNews); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the repository to update the news
	news, err := ac.repo.Update(c, uint(id), updatedNews)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated news
	c.JSON(http.StatusOK, news)
}

// DeleteNews handles the HTTP request to delete an news by its ID.
func (ac *NewsController) DeleteNews(c *gin.Context) {
	// Extracting the news ID from the route parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
		return
	}

	// Call the Delete method of the repository to delete the news
	if err := ac.repo.Delete(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting news"})
		return
	}

	// If the deletion is successful, return a no content response
	c.JSON(http.StatusNoContent, nil)
}
