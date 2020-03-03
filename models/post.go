package models

// 帖子
type Post struct {
	Model
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
	Summary string `json:"summary" gorm:"column:summary"`
	UserId  uint   `json:"user_id" gorm:"column:user_id"`
}
