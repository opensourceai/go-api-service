package models

// 帖子
type Post struct {
	Model
	Title   string `json:"title" gorm:"column:title" valid:"Required;Min(1)"`           // 标题
	Content string `json:"content" gorm:"column:content" valid:"Required;"`             // 内容
	Summary string `json:"summary" gorm:"column:summary" valid:"Required;MaxSize(200)"` // 摘要
	UserId  int    `json:"user_id" gorm:"column:user_id" valid:"Required;"`             // 用户ID
}
