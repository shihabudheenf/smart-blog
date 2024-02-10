package dbutil // define a new package for database utilities

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetDBFromContext retrieves a *gorm.DB instance from the Gin context.
// It returns an error if the connection does not exist or is of the wrong type.
func GetDBFromContext(c *gin.Context) (*gorm.DB, error) {
	db, exists := c.Get("db")
	if !exists {
		return nil, fmt.Errorf("database connection not found in context")
	}

	gormDB, ok := db.(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("context contains a non-GORM database connection")
	}

	return gormDB, nil
}
