package main

import (
	"log"
	"smart-blog/database"
	articleRoutes "smart-blog/blog/articles/routes" // Alias for articles routes package
    newsRoutes "smart-blog/blog/news/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.GetDatabaseInstance()

	r := gin.Default()

	// Middleware to set db in context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	articleRoutes.RegisterArticleRoutes(r)
	newsRoutes.RegisterNewsRoutes(r)

	log.Println("Connected to the database successfully!", db)

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}

}
