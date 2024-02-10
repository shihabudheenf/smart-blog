package controllers

import (
	"net/http"
	"smart-blog/blog/articles/models"
	"smart-blog/blog/articles/repositories/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	repo interfaces.ArticleRepositoryInterface
}

func NewArticleController(repo interfaces.ArticleRepositoryInterface) *ArticleController {
	return &ArticleController{
		repo: repo,
	}
}

func (ac *ArticleController) GetAllArticles(c *gin.Context) {
	articles, err := ac.repo.GetAll(c) // Pass the Gin context here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (ac *ArticleController) GetArticleByID(c *gin.Context) {

	// Extract the 'id' path parameter and convert it to an unsigned integer
	idStr := c.Param("id")                      // 'id' is the name of the path parameter
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID format"})
		return
	}

	article, err := ac.repo.GetByID(c, uint(id)) // Note the conversion of id to uint
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

// CreateArticle handles the POST request for creating a new article
func (ac *ArticleController) CreateArticle(c *gin.Context) {
	var article models.Article

	// Bind the JSON payload to the Article struct
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the repository's Create method, passing the bound article
	_, err := ac.repo.Create(c, article)
	if err != nil {
		// Handle errors that may occur during article creation
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If successful, respond with the created article
	c.Status(http.StatusCreated)

}

func (ac *ArticleController) UpdateArticle(c *gin.Context) {
	// Assuming "id" is a parameter in the route
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid article ID"})
		return
	}

	// Bind the request body to the Article struct
	var updatedArticle models.ArticleUpdateInput
	if err := c.ShouldBindJSON(&updatedArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the repository to update the article
	article, err := ac.repo.Update(c, uint(id), updatedArticle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated article
	c.JSON(http.StatusOK, article)
}

// DeleteArticle handles the HTTP request to delete an article by its ID.
func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	// Extracting the article ID from the route parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	// Call the Delete method of the repository to delete the article
	if err := ac.repo.Delete(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting article"})
		return
	}

	// If the deletion is successful, return a no content response
	c.JSON(http.StatusNoContent, nil)
}
