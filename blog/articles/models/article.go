package models

type Article struct {
	ID         int      `gorm:"primary_key" json:"id"`
	CategoryID int      `json:"category_id" binding:"required,gt=0" gorm:"not null"`
	Title      string   `json:"title" binding:"required,max=500"`
	Author     string   `json:"author" binding:"required,max=200"`
	Category   Category `gorm:"foreignkey:CategoryID" json:"category"`
}

type ArticleUpdateInput struct {
	CategoryID *int    `json:"category_id,omitempty"`
	Title      *string `json:"title,omitempty"`
	Author     *string `json:"author,omitempty"`
}
