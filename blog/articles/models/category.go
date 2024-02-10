package models

// Category represents the category model that matches the 'categories' table in the database.
type Category struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
}
