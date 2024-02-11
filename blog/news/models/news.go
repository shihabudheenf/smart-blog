package models

type News struct {
	ID         int      `gorm:"primary_key" json:"id"`
	Title      string   `json:"title" binding:"required,max=500"`
	Content     string   `json:"content" binding:"required"`
}

type NewsUpdateInput struct {
	Title      *string `json:"title,omitempty"`
	Content     *string `json:"content,omitempty"`
}
